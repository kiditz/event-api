package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/labstack/echo/v4"
)

// AddToCart godoc
func AddToCart(cart *e.Cart, c echo.Context) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		cart.IPAddress = c.RealIP()
		if err := tx.Create(&cart).Error; err != nil {
			return err
		}
		return nil
	})
}

//DeleteCart delete cart by loggedin
func DeleteCart(c echo.Context) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		ipAddress := c.RealIP()
		if err := tx.Where("ip_address = ?", ipAddress).Delete(e.Cart{}).Error; err != nil {
			return err
		}
		return nil
	})
}
