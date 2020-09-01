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
	db = db.Where("is_searchable = ?", true).Order("name asc").Find(&subCategories)
	return subCategories
}

// GetExpertises docs
func GetExpertises() []e.Expertise {
	var expertises []e.Expertise
	db.DB.Select("distinct ON (name) name, slug, id").Find(&expertises)
	return expertises
}

// GetBusinesType docs
func GetBusinesType() []e.BusinessType {
	var businessTypes []e.BusinessType
	db.DB.Select("distinct ON (name) name, slug, id").Find(&businessTypes)
	return businessTypes
}

// GetOccupations docs
func GetOccupations() []e.Occupation {
	var occupations []e.Occupation
	db.DB.Select("distinct ON (name) name, slug, id").Find(&occupations)
	return occupations
}

// GetSalaryRates docs
func GetSalaryRates(subCategoryID int) []e.Rate {
	var rates []e.Rate
	db.DB.Joins("JOIN sub_categories s ON rates.sub_category_slug = s.slug").Where("s.id = ?", subCategoryID).Order("price desc").Find(&rates)
	return rates
}
