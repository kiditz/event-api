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
		&e.DigitalStaff{},
		&e.EventStaff{},
		&e.Image{},
		&e.SocialMedia{},
		&e.Document{},
	)
	db.DB.Model(&e.Campaign{}).AddUniqueIndex("idx_created_by_title", "created_by", "title")
}
