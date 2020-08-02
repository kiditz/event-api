package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddEventStaff event staff category
func AddEventStaff(eventStaff *e.EventStaff) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&eventStaff).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetEventStaff get event staff category
func GetEventStaff() ([]e.EventStaff, error) {
	var records []e.EventStaff
	if err := db.DB.Find(&records).Order("id", false).Error; err != nil {
		return records, err
	}
	return records, nil
}
