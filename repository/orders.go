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
		if err := tx.Create(&cart).Error; err != nil {
			return err
		}
		return nil
	})
}

//DeleteCart delete cart by loggedin
func DeleteCart(deviceID string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("device_id = ?", deviceID).Delete(e.Cart{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//GetCarts delete cart by loggedin
func GetCarts(deviceID string) []e.Cart {
	carts := []e.Cart{}
	db.DB.Where("device_id = ?", deviceID).Preload("Service.Category").Preload("Service.SubCategory").Preload("Talent.User").Find(&carts)
	return carts
}

// AddInvitation godoc
func AddInvitation(invitations *[]e.Invitation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		for _, invitation := range *invitations {
			if err := tx.Create(&invitation).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

//GetInvitations find invitations by talent logged in
func GetInvitations(email string, limitOffset e.LimitOffset) []e.Invitation {
	if limitOffset.Limit <= 0 {
		limitOffset.Limit = 10
	}
	invitations := []e.Invitation{}
	query := db.DB
	query = query.Joins("JOIN services s ON s.id = invitations.service_id")
	query = query.Joins("JOIN talents t ON t.id = s.talent_id")
	query = query.Joins("JOIN users u ON t.user_id = u.id")
	query = query.Where("u.email = ?", email)
	query = query.Preload("Campaign")
	query = query.Preload("Service.Category").Preload("Service.SubCategory").Preload("Service.Topic")
	query = query.Preload("Campaign.Company.Image")
	query = query.Preload("Campaign.Location")
	query = query.Preload("Campaign.Company")

	query = query.Order("id desc").Limit(limitOffset.Limit).Offset(limitOffset.Offset).Find(&invitations)
	return invitations
}

// AddQuotation godoc
func AddQuotation(quote *e.Quotation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&quote).Error; err != nil {
			return err
		}
		return nil
	})
}
