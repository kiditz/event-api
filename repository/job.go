package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-job/db"
	e "github.com/kiditz/spgku-job/entity"
)

const (
	// StatusClosed job status is closed
	StatusClosed = "CLOSED"
	// StatusOpen job status is open
	StatusOpen = "OPEN"
)

// CreateJob used to create new job operation
func CreateJob(event *e.Job) error {

	return db.DB.Transaction(func(tx *gorm.DB) error {
		event.Status = StatusOpen
		if err := tx.Create(&event).Error; err != nil {
			return err
		}
		return nil
	})
}
