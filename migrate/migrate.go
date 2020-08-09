package main

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

func main() {
	db.Connect()
	defer db.DB.Close()
	db.DB.AutoMigrate(
		&e.Campaign{},
		&e.Location{},
		&e.User{},
		&e.SubCategory{},
		&e.Category{},
		&e.Image{},
		&e.SocialMedia{},
		&e.Document{},
		&e.SubCategory{},
		&e.Talent{},
		&e.Expertise{},
		&e.Service{},
	)
	db.DB.Model(&e.Campaign{}).AddUniqueIndex("idx_created_by_title", "created_by", "title")
}
