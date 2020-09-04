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
func AddBrief(campaign *e.Brief) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false)
		if err := tx.Save(&campaign).Error; err != nil {
			return err
		}

		if campaign.Location != nil {
			if err := tx.Where("formatted_address = ?", campaign.Location.FormattedAddress).First(campaign.Location.FormattedAddress).First(&campaign.Location).Error; err != nil {
				tx.Model(&campaign.Location).Save(campaign.Location)
			}
			tx.Model(&campaign).Association("Location").Append(campaign.Location)
		}
		if campaign.PaymentTerms != nil {
			tx.Model(&campaign.PaymentTerms).Save(campaign.PaymentTerms)
			tx.Model(&campaign).Association("PaymentTerms").Append(campaign.PaymentTerms)
		}
		if campaign.PaymentDays != nil {
			tx.Model(&campaign.PaymentDays).Save(campaign.PaymentDays)
			tx.Model(&campaign).Association("PaymentDays").Append(campaign.PaymentDays)
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
	if len(filter.Date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.Date)
	}
	if len(filter.Query) > 0 {
		query = query.Where("title ilike ?", "%"+filter.Query+"%")
	}
	if filter.OnlyMe {
		query = query.Where("created_by = ?", utils.GetEmail(c))
	}
	query = query.Preload("Company.Image").Preload("Company.BackgroundImage")
	query = query.Preload("Location")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc").Find(&campaign)
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
