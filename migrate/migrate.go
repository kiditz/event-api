package main

import (
	"github.com/kiditz/spgku-job/db"
	e "github.com/kiditz/spgku-job/entity"
)

func main() {
	db.Connect()
	defer db.DB.Close()
	db.DB.AutoMigrate(
		&e.Job{},
		&e.Company{},
	)
	db.DB.Model(&e.Job{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
}
