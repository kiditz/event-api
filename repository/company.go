package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// FindCompany docs
func FindCompany(c echo.Context) (e.Company, error) {
	var company e.Company
	user := utils.GetUser(c)

	userID := uint(user["id"].(float64))

	if err := db.DB.Where("user_id = ?", userID).Preload("Image").Find(&company).Error; err != nil {
		return company, err
	}
	return company, nil
}

// UpdateCompany godoc
func UpdateCompany(company *e.Company, c echo.Context) error {
	user := utils.GetUser(c)
	company.UserID = uint(user["id"].(float64))
	company.CreatedBy = user["email"].(string)
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&company).Error; err != nil {
			return err
		}
		return nil
	})
}
