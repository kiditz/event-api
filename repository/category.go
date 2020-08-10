package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

//GetCategories find all categories
func GetCategories() []e.Category {
	var categories []e.Category
	db.DB.Find(&categories)
	return categories
}

//GetSubCategories find all sub_categories
func GetSubCategories() []e.SubCategory {
	var subCategories []e.SubCategory
	db.DB.Where("is_searchable = ?", true).Find(&subCategories)
	return subCategories
}

//GetSubCategoriesByCategoryID find all sub_categories
func GetSubCategoriesByCategoryID(categoryID int) []e.SubCategory {
	var subCategories []e.SubCategory
	db.DB.Where("category_id = ?", categoryID).Where("is_searchable = ?", true).Find(&subCategories)
	return subCategories
}
