package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-job/db"
	e "github.com/kiditz/spgku-job/entity"
)

// CreateCompany used for inserting company into "companies" database "params" companies id required
func CreateCompany(company *e.Company) error {

	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&company).Error; err != nil {
			return err
		}
		return nil
	})
}
