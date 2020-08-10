package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddCampaign used to insert campaign into campaigns database
func AddCampaign(campaign *e.Campaign) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false)
		if err := tx.Save(&campaign).Error; err != nil {
			return err
		}
		if len(campaign.Services) > 0 {
			tx.Model(&campaign.Services).Association("Services").Append(campaign.Services)
		}
		if campaign.Location != nil {
			tx.Model(&campaign.Location).Save(campaign.Location)
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

// FindCampaignByID  used to find campaign by id
func FindCampaignByID(campaignID int) (e.Campaign, error) {
	var campaign e.Campaign
	if err := db.DB.Where("id=?", campaignID).Preload("Location").Preload("PaymentTerms").Preload("PaymentDays").Find(&campaign).Error; err != nil {
		return campaign, err
	}
	return campaign, nil
}

// CampaignsFilter filtering query by
type CampaignsFilter struct {
	Date   string `query:"date" json:"date"`
	Query  string `query:"q" json:"q"`
	Offset int    `query:"offset" json:"offset"`
	Limit  int    `query:"limit" json:"limit"`
	OnlyMe bool   `query:"onlyme" json:"onlyme"`
}

// GetCampaigns used to find campaign by filtered values
func GetCampaigns(filter *CampaignsFilter, c echo.Context) []e.Campaign {
	var campaign []e.Campaign
	query := db.DB
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	fmt.Printf("Filtered :%v", &filter)
	if len(filter.Date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.Date)
	}
	if len(filter.Query) > 0 {
		query = query.Where("title ilike ?", "%"+filter.Query+"%")
	}
	if filter.OnlyMe {
		query = query.Where("created_by = ?", utils.GetEmail(c))
	}

	query = query.Preload("PaymentTerms").Preload("PaymentDays").Offset(filter.Offset).Limit(filter.Limit).Order("id desc").Find(&campaign)
	return campaign
}

// GetAllSocialMedia docs
func GetAllSocialMedia() []e.SocialMedia {
	var socialMediaList []e.SocialMedia
	db.DB.Find(&socialMediaList)
	return socialMediaList
}

// GetPaymentTerms docs
func GetPaymentTerms() []e.PaymentTerms {
	var paymentTerms []e.PaymentTerms
	db.DB.Find(&paymentTerms)
	return paymentTerms
}

// GetPaymentDays docs
func GetPaymentDays() []e.PaymentDays {
	var paymentDays []e.PaymentDays
	db.DB.Find(&paymentDays)
	return paymentDays
}
