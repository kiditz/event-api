package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
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
type BriefsFilter struct {
	Date   string `query:"date" json:"date"`
	Query  string `query:"q" json:"q"`
	Offset int    `query:"offset" json:"offset"`
	Limit  int    `query:"limit" json:"limit"`
	OnlyMe bool   `query:"onlyme" json:"onlyme"`
}

// GetBriefs used to find campaign by filtered values
func GetBriefs(filter *BriefsFilter, c echo.Context) []e.Brief {
	var campaign []e.Brief
	query := db.DB
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	email := utils.GetEmail(c)
	if len(filter.Date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.Date)
	}
	if len(filter.Query) > 0 {
		query = query.Where("title ilike ?", "%"+filter.Query+"%")
	}
	if filter.OnlyMe {
		query = query.Where("created_by = ?", email)
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
