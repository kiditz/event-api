package main

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

func main() {
	db.Connect()
	defer db.DB.Close()
	db.DB.AutoMigrate(
		&e.BusinessType{},
		&e.Brief{},
		&e.Cart{},
		&e.Category{},
		&e.Company{},
		&e.Document{},
		&e.Expertise{},
		&e.Image{},
		&e.Invitation{},
		&e.Location{},
		&e.Occupation{},
		&e.PaymentDays{},
		&e.PaymentTerms{},
		&e.Portfolio{},
		&e.Quotation{},
		&e.Rate{},
		&e.Service{},
		&e.SocialMedia{},
		&e.User{},
		&e.SubCategory{},
		&e.Talent{},
		&e.Order{},
		&e.TransactionDetails{},
		&e.ItemDetails{},
		&e.PaymentNotification{},
		&e.Income{},
		&e.Billing{},
	)
	db.DB.Model(&e.Invitation{}).AddUniqueIndex("uk_invitation_service_id_brief_id", "service_id", "brief_id")
	db.DB.Model(&e.Income{}).AddUniqueIndex("uk_income_user_id_brief_id", "user_id", "brief_id")
	db.DB.Model(&e.Billing{}).AddUniqueIndex("uk_billing_user_id_brief_id", "user_id", "brief_id")
}
