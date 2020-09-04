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
	)
	// db.DB.Model(&e.Cart{}).AddUniqueIndex("cart_service_id_device_id", "service_id", "device_id")
	// db.DB.Model(&e.Cart{}).AddUniqueIndex("cart_category_id_device_id", "category_id", "device_id")
}
