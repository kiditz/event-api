package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddCampaign used to insert job into campaign into campaigns database
func AddCampaign(campaign *e.Campaign) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Set("gorm:association_autocreate", false).Save(&campaign).Error; err != nil {
			return err
		}
		tx.Model(&campaign).Association("SocialMedias").Append(campaign.SocialMedias)
		return nil
	})
}
