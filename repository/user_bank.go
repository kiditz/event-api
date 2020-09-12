package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// GetBanks godoc
func GetBanks() []e.Bank {
	var banks []e.Bank
	db.DB.Find(&banks)
	return banks
}

// AddUserBank godoc
func AddUserBank(c echo.Context, userBank *e.UserBank) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		user := utils.GetUser(c)
		userID := uint(user["id"].(float64))
		userBank.UserID = userID
		if err := tx.Save(&userBank).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetUserBank godoc
func GetUserBank(c echo.Context) []e.UserBank {
	var userBanks []e.UserBank
	user := utils.GetUser(c)
	userID := uint(user["id"].(float64))
	db.DB.Where("user_id = ?", userID).Preload("Bank").Find(&userBanks)
	return userBanks
}
