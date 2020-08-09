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
		tx.Model(&campaign.SocialMedias).Association("SocialMedias").Append(campaign.SocialMedias)
		if campaign.Location != nil {
			tx.Model(&campaign.Location).Save(campaign.Location)
			tx.Model(&campaign).Association("Location").Append(campaign.Location)
		}
		if campaign.Images != nil {
			tx.Model(&campaign.Images).Association("Images").Append(campaign.Images)
		}
		return nil
	})
}

// FindCampaignByID  used to find campaign by id
func FindCampaignByID(campaignID int) (e.Campaign, error) {
	var campaign e.Campaign
	if err := db.DB.Where("id=?", campaignID).Preload("Location").Preload("Images").Preload("SocialMedias").Find(&campaign).Error; err != nil {
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

// GetCampaigns  used to find campaign by date
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

	query = query.Preload("Images").Offset(filter.Offset).Limit(filter.Limit).Order("id desc").Find(&campaign)
	return campaign
}

// GetAllSocialMedia  used to find campaign by date
func GetAllSocialMedia() ([]e.SocialMedia, error) {
	var socialMediaList []e.SocialMedia
	if err := db.DB.Find(&socialMediaList).Error; err != nil {
		return socialMediaList, err
	}
	return socialMediaList, nil
}
