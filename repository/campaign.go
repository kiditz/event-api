package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
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
	date   string `query:"name"`
	q      string `query:"q"`
	offset int    `query:"offset"`
	limit  int    `query:"limit"`
}

// GetCampaigns  used to find campaign by date
func GetCampaigns(filter *CampaignsFilter) []e.Campaign {
	var campaign []e.Campaign
	query := db.DB
	if filter.limit == 0 {
		filter.limit = 10
	}
	if len(filter.date) > 0 {
		query = query.Where("? between to_char(start_date, 'YYYY-MM-DD') and to_char(end_date, 'YYYY-MM-DD')", filter.date)
	}
	if len(filter.q) > 0 {
		query = query.Where("title LIKE ?", "%"+filter.q+"%")
	}

	query = query.Offset(filter.offset).Limit(filter.limit).Order("id desc").Find(&campaign)
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
