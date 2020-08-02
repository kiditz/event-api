package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddDigitalStaff used for inserting company into "companies" database "params" companies id required
func AddDigitalStaff(digitalStaff *e.DigitalStaff) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&digitalStaff).Error; err != nil {
			return err
		}
		return nil
	})
}
