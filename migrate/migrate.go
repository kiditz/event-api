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
		&e.Category{},
		&e.Document{},
		&e.Expertise{},
		&e.Image{},
		&e.Location{},
		&e.PaymentDays{},
		&e.PaymentTerms{},
		&e.Service{},
		&e.SocialMedia{},
		&e.User{},
		&e.SubCategory{},
		&e.Talent{},
	)
	db.DB.Model(&e.Campaign{}).AddUniqueIndex("idx_created_by_title", "created_by", "title")
}
