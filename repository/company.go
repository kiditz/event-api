package repository

import (
	"fmt"

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

// FindCompanyByID is used to query company by it's primary ke
func FindCompanyByID(companyID int) (e.Company, error) {
	var company e.Company
	if db.DB.Find(&company, companyID).RecordNotFound() {
		err := fmt.Errorf("company_not_found")
		return company, err
	}
	return company, nil
}
