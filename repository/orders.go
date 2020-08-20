package repository

import (
	"fmt"

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
			invitation.Status = e.ACTIVE
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
	query = query.Preload("Campaign.SubCategory")
	query = query.Preload("Campaign.Category")
	query = query.Preload("Campaign.Location")
	query = query.Preload("Campaign.Company")
	query = query.Preload("Campaign.PaymentTerms")
	query = query.Order("id desc").Limit(limitOffset.Limit).Offset(limitOffset.Offset).Find(&invitations)
	return invitations
}

// AcceptInvitation godoc
func AcceptInvitation(quote *e.Quotation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var invitation e.Invitation
		if err := tx.Where("id = ?", quote.InvitationID).Find(&invitation).Error; err != nil {
			return err
		}
		if invitation.Status == e.ACCEPTED {
			return fmt.Errorf("status_was_accepted")
		}
		// Save Invitation
		invitation.Status = e.ACCEPTED
		if err := tx.Save(&invitation).Error; err != nil {
			return err
		}
		// Save Quotation
		quote.Status = "active"
		if err := tx.Save(&quote).Error; err != nil {
			return err
		}
		return nil
	})
}

// RejectInvitation godoc
func RejectInvitation(reject *e.RejectInvitation) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		var invitation e.Invitation
		if err := tx.Where("id = ?", reject.InvitationID).Find(&invitation).Error; err != nil {
			return err
		}
		if invitation.Status == e.ACCEPTED {
			return fmt.Errorf("reject_is_not_allowed")
		}
		// Save Invitation
		invitation.Status = e.REJECTED
		if err := tx.Save(&invitation).Error; err != nil {
			return err
		}
		return nil
	})
}

//GetQuotations used to find quotations list by campaig id
func GetQuotations(filter e.FilteredQuotations) []e.QuotationList {
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	quotations := []e.QuotationList{}
	query := db.DB.Table("quotations q")
	query = query.Select("q.id, q.price, q.message, u.name, i.image_url, q.status")
	query = query.Joins("JOIN services s ON q.service_id = s.id")
	query = query.Joins("JOIN talents t ON t.id = s.talent_id")
	query = query.Joins("JOIN users u ON t.user_id = u.id")
	query = query.Joins("JOIN images i ON t.image_id = i.id")
	query = query.Where("q.campaign_id = ?", filter.CampaignID)
	rows, err := query.Offset(filter.Offset).Limit(filter.Limit).Order("q.id desc").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Printf("Wrong query :%v", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		quotation := e.QuotationList{}
		err := rows.Scan(&quotation.ID, &quotation.Price, &quotation.Message, &quotation.Name, &quotation.ImageURL, &quotation.Status)
		if err != nil {
			fmt.Println(err)
		}
		quotations = append(quotations, quotation)
	}
	return quotations
}
