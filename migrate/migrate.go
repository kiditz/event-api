package main

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

func main() {
	db.Connect()
	defer db.DB.Close()
	db.DB.AutoMigrate(
		&e.Job{},
		&e.Company{},
		&e.Location{},
		&e.User{},
		&e.DigitalStaff{},
		&e.EventStaff{},
	)
	db.DB.Model(&e.Job{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	db.DB.Model(&e.Company{}).AddForeignKey("location_id", "locations(id)", "CASCADE", "CASCADE")
	db.DB.Model(&e.Job{}).AddUniqueIndex("idx_company_id_title", "company_id", "title", "status")
}
