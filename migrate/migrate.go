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
		&e.Campaign{},
		&e.Cart{},
		&e.Category{},
		&e.Company{},
		&e.Document{},
		&e.Expertise{},
		&e.Image{},
		&e.Invitation{},
		&e.Location{},
		&e.PaymentDays{},
		&e.PaymentTerms{},
		&e.Quotation{},
		&e.Service{},
		&e.SocialMedia{},
		&e.User{},
		&e.SubCategory{},
		&e.Talent{},
	)
	db.DB.Model(&e.Campaign{}).AddUniqueIndex("idx_created_by_title", "created_by", "title")
	db.DB.Model(&e.Talent{}).AddUniqueIndex("talent_idx_created_by", "created_by")
	db.DB.Model(&e.Cart{}).AddUniqueIndex("cart_service_id_talent_id", "service_id", "talent_id")
}
