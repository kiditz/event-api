package repository

import (
	"fmt"
	"time"

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

// FindCampaignByID  used to find campaign by id
func FindCampaignByID(campaignID int) (e.Campaign, error) {
	var campaign e.Campaign
	if err := db.DB.Where("id=?", campaignID).Preload("Category").Preload("SubCategory").Preload("Location").Preload("PaymentTerms").Preload("PaymentDays").Find(&campaign).Error; err != nil {
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
	if len(filter.Date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.Date)
	}
	if len(filter.Query) > 0 {
		query = query.Where("title ilike ?", "%"+filter.Query+"%")
	}
	if filter.OnlyMe {
		query = query.Where("created_by = ?", utils.GetEmail(c))
	}
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

// GetCampaignInfo godoc
func GetCampaignInfo(campaignID int) (e.CampaignInfo, error) {
	info := e.CampaignInfo{}
	campaign := e.Campaign{}
	query := db.DB
	if err := query.Where("id = ?", campaignID).Find(&campaign).Error; err != nil {
		return info, err
	}
	query.Model(e.Quotation{}).Where("campaign_id = ?", campaignID).Count(&info.ApprovedCount)
	info.StaffAmount = campaign.StaffAmount
	query = query.Table("quotations q")
	query = query.Select("i.image_url")
	query = query.Joins("JOIN services s ON s.id = q.service_id")
	query = query.Joins("JOIN talents t ON t.id = s.talent_id")
	query = query.Joins("JOIN images i ON i.id = t.image_id")
	rows, _ := query.Where("q.campaign_id = ? AND status = ?", campaignID, e.APPROVED).Order("q.id desc").Limit(5).Rows()
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

// StartCampiagn godoc
func StartCampiagn(campaignID int) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var campaign e.Campaign
		if err := tx.Find(campaign, campaignID).Error; err != nil {
			return err
		}
		now := time.Now()
		campaign.StartDate = &now
		tx.Save(campaign)
		return nil
	})
}

// StopProject godoc
func StopProject(campaignID int) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var campaign e.Campaign
		if err := tx.Find(campaign, campaignID).Error; err != nil {
			return err
		}
		now := time.Now()
		campaign.EndDate = &now
		tx.Save(campaign)
		return nil
	})
}
