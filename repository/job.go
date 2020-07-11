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

// CreateJob used to insert job into jobs database
func CreateJob(job *e.Job) error {

	return db.DB.Transaction(func(tx *gorm.DB) error {
		job.Status = StatusOpen
		if err := tx.Create(&job).Error; err != nil {
			return err
		}
		return nil
	})
}
