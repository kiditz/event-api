package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

//GetCategories find all categories
func GetCategories() []e.Category {
	var categories []e.Category
	db.DB.Preload("SubCategories").Find(&categories)
	return categories
}
