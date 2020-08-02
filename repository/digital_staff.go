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

// GetDigitalStaff used for inserting company into "companies" database "params" companies id required
func GetDigitalStaff() ([]e.DigitalStaff, error) {
	var records []e.DigitalStaff
	if err := db.DB.Find(&records).Order("id", false).Error; err != nil {
		return records, err
	}
	return records, nil
}
