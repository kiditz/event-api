package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	t "github.com/kiditz/spgku-api/translate"
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
	if db.DB.Preload("Location").Find(&company, companyID).RecordNotFound() {
		err := fmt.Errorf("company_not_found")
		return company, err
	}
	return company, nil
}

// GetCompanies by input.Name or input.Location if exists
func GetCompanies(name string, country string, city string, offset int, limit int) []e.Company {
	if limit == 0 {
		limit = 10
	}

	var companies []e.Company
	tx := db.DB.Model(&companies).Preload("Location")
	if t.NotEmpty(name) {
		tx = db.DB.Where("name LIKE ?", "%"+name+"%")
	}

	if t.NotEmpty(country) {
		tx = tx.Joins("inner join locations ON locations.id = companies.location_id AND locations.country LIKE ?", "%"+country+"%")
	}

	if t.NotEmpty(city) {
		tx = tx.Joins("inner join locations ON locations.id = companies.location_id AND locations.city LIKE ?", "%"+city+"%")
	}

	tx = tx.Offset(offset).Limit(limit)
	tx.Find(&companies)
	return companies
}
