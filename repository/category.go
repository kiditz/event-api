package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

//GetCategories find all categories
func GetCategories() []e.Category {
	var categories []e.Category
	db.DB.Order("id").Find(&categories)
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
	db := db.DB
	if categoryID > 1 {
		db = db.Where("category_id = ?", categoryID)
	}
	db = db.Where("is_searchable = ?", true).Find(&subCategories)
	return subCategories
}

// GetExpertises docs
func GetExpertises() []e.Expertise {
	var expertises []e.Expertise
	db.DB.Find(&expertises)
	return expertises
}
