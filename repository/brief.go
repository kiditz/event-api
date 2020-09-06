package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddBrief used to insert campaign into briefs database
func AddBrief(c echo.Context, brief *e.Brief) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		company, err := FindCompany(c)
		if err != nil {
			return err
		}
		brief.CompanyID = company.ID
		brief.CreatedBy = utils.GetEmail(c)
		tx = tx.Set("gorm:association_autoupdate", false)
		if err := tx.Save(&brief).Error; err != nil {
			return err
		}

		if brief.Location != nil {
			if err := tx.Where("formatted_address = ?", brief.Location.FormattedAddress).First(brief.Location.FormattedAddress).First(&brief.Location).Error; err != nil {
				tx.Model(&brief.Location).Save(brief.Location)
			}
			tx.Model(&brief).Association("Location").Append(brief.Location)
		}
		if brief.PaymentTerms != nil {
			tx.Model(&brief.PaymentTerms).Save(brief.PaymentTerms)
			tx.Model(&brief).Association("PaymentTerms").Append(brief.PaymentTerms)
		}
		if brief.PaymentDays != nil {
			tx.Model(&brief.PaymentDays).Save(brief.PaymentDays)
			tx.Model(&brief).Association("PaymentDays").Append(brief.PaymentDays)
		}
		return nil
	})
}

// FindBriefByID  used to find campaign by id
func FindBriefByID(campaignID int) (e.Brief, error) {
	var campaign e.Brief
	if err := db.DB.Where("id=?", campaignID).Preload("Company.BackgroundImage").Preload("Company.Image").Preload("Category").Preload("Location").Preload("PaymentTerms").Preload("PaymentDays").Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

// BriefsFilter filtering query by

// GetBriefs used to find campaign by filtered values
func GetBriefs(filter *e.BriefsFilter, c echo.Context) []e.Brief {
	var campaign []e.Brief
	query := db.DB
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	email := utils.GetEmail(c)
	user := utils.GetUser(c)
	userType := user["type"].(string)
	if len(filter.Date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.Date)
	}
	if len(filter.Query) > 0 {
		query = query.Where("title ilike ?", "%"+filter.Query+"%")
	}
	if filter.OnlyMe {
		if userType == "company" {
			query = query.Where("created_by = ?", email)

		} else {
			query = query.Joins("JOIN quotations q ON briefs.id =  q.brief_id ")
			query = query.Joins("JOIN services s ON s.id = q.service_id")
			query = query.Joins("JOIN users u ON u.id = s.user_id AND u.email = ?", email)
		}

	} else {
		query = query.Joins("JOIN users u ON u.email = ? ", email)
		query = query.Joins("JOIN services s ON s.user_id = u.id AND s.category_id = briefs.category_id")
		query = query.Where("briefs.status = ?", e.BOOKING)
		query = query.Where(`
			NOT EXISTS (
				SELECT 1 FROM quotations q 
				JOIN services s ON q.service_id = s.id 
				JOIN users u ON u.id = s.user_id 
				WHERE u.email = ?
				AND q.brief_id = briefs.id AND q.status = 'approved'
			)`, email)
	}
	query = query.Preload("Company.Image").Preload("Company.BackgroundImage")
	query = query.Preload("Location")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc").Select("DISTINCT briefs.*").Find(&campaign)
	return campaign
}

// GetAllSocialMedia godoc
func GetAllSocialMedia() []e.SocialMedia {
	var socialMediaList []e.SocialMedia
	db.DB.Find(&socialMediaList)
	return socialMediaList
}

// GetPaymentTerms godoc
func GetPaymentTerms() []e.PaymentTerms {
	var paymentTerms []e.PaymentTerms
	db.DB.Find(&paymentTerms)
	return paymentTerms
}

// GetPaymentDays godoc
func GetPaymentDays() []e.PaymentDays {
	var paymentDays []e.PaymentDays
	db.DB.Find(&paymentDays)
	return paymentDays
}

// GetBriefInfo godoc
func GetBriefInfo(campaignID int) (e.BriefInfo, error) {
	info := e.BriefInfo{Images: []string{}}
	campaign := e.Brief{}
	query := db.DB
	if err := query.Where("id = ?", campaignID).Find(&campaign).Error; err != nil {
		return info, err
	}
	query.Model(e.Quotation{}).Where("brief_id = ? and status = ?", campaignID, e.APPROVED).Count(&info.ApprovedCount)
	query.Model(e.Quotation{}).Where("brief_id = ?", campaignID).Count(&info.QuotationCount)
	info.StaffAmount = campaign.StaffAmount
	query = query.Table("quotations q")
	query = query.Select("u.image_url")
	query = query.Joins("JOIN services s ON s.id = q.service_id")
	query = query.Joins("JOIN users u ON u.id = s.user_id")
	rows, _ := query.Where("q.brief_id = ? AND q.status = ?", campaignID, e.APPROVED).Order("q.id desc").Limit(5).Rows()
	defer rows.Close()
	for rows.Next() {
		var image string
		err := rows.Scan(&image)
		if err != nil {
			fmt.Println(err)
		}
		info.Images = append(info.Images, image)
	}

	return info, nil
}

//StopBrief stop campaign
func StopBrief(c echo.Context, briefID uint) error {

	return db.DB.Transaction(func(tx *gorm.DB) error {

		var brief e.Brief
		if tx.Where("status = ? AND id = ?", e.ACTIVE, briefID).Preload("PaymentTerms").Preload("PaymentDays").First(&brief).RecordNotFound() {
			return fmt.Errorf("brief_not_found")
		}
		var order e.Order
		query := tx
		query = query.Preload("TransactionDetails")
		query = query.Preload("ItemDetails")
		if query.Where("brief_id = ?", briefID).Last(&order).RecordNotFound() {
			return fmt.Errorf("order_not_found")
		}
		user := utils.GetUser(c)
		userID := uint(user["id"].(float64))
		now := time.Now().UTC()
		// Close the brief
		brief.Status = e.CLOSED
		if err := tx.Save(&brief).Error; err != nil {
			return err
		}
		if brief.PaymentTerms.Slug == front {
			if order.TransactionStatus == txCapture || order.TransactionStatus == txSettlement {
				return addIncome(tx, &order, &brief, true)
			}
		} else if brief.PaymentTerms.Slug == back || brief.PaymentTerms.Slug == fiftyFifty {
			dueDate := now.AddDate(0, 0, int(brief.PaymentDays.Days))
			billing := e.Billing{
				BriefID: brief.ID,
				Amount:  order.TransactionDetails.Billing,
				DueDate: &dueDate,
				OrderID: order.TransactionDetails.OrderID,
				UserID:  userID,
			}
			if err := tx.Save(&billing).Error; err != nil {
				return err
			}
			return addIncome(tx, &order, &brief, false)
		}

		fmt.Println("stop brief")
		return nil
	})
}
func addIncome(tx *gorm.DB, order *e.Order, brief *e.Brief, canWithdrawal bool) error {
	for _, item := range order.ItemDetails {
		if !translate.IsEmail(item.Name) {
			continue
		}
		fmt.Printf("Item %v\n", item)
		user, err := FindUserByEmail(item.Name)
		if err != nil {
			fmt.Printf("user %s not found", item.Name)
			continue
		}
		income := e.Income{
			BriefID:       brief.ID,
			Amount:        item.Price,
			UserID:        user.ID,
			OrderID:       order.TransactionDetails.OrderID,
			CanWithdrawal: true,
			HasWithdraw:   false,
		}
		if err := tx.Save(&income).Error; err != nil {
			return err
		}
	}
	return nil
}

const (
	fiftyFifty = "50_50"
	front      = "100_front"
	back       = "100_back"
)
