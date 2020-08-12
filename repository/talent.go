package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddTalent used to create new talent
func AddTalent(talent *e.Talent) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&talent).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindTalentByID  used to find talent by id
func FindTalentByID(talentID int) (e.Talent, error) {
	var talent e.Talent
	if err := db.DB.Where("id=?", talentID).Preload("Location").Preload("Image").Find(&talent).Error; err != nil {
		return talent, err
	}
	return talent, nil
}

// FilteredTalent used to filter talent query
type FilteredTalent struct {
	CategoryID    int64  `query:"category_id" json:"category_id"`
	ExpertiseName string `query:"expertise_name" json:"expertise_name"`
	Limit         int64  `query:"limit" json:"limit"`
	Offset        int64  `query:"offset" json:"offset"`
	Q             string `query:"q" json:"q"`
	SubCategoryID int64  `query:"sub_category_id" json:"sub_category_id"`
}

// GetTalents  used to find talent by id
func GetTalents(filter *FilteredTalent) []e.Talent {
	var talents []e.Talent
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	query := db.DB.Joins("JOIN services s ON s.talent_id = talents.id")
	if filter.CategoryID > 1 {
		query = query.Where("s.category_id = ?", filter.CategoryID)
	}
	if filter.SubCategoryID > 0 {
		query = query.Where("s.sub_category_id = ?", filter.SubCategoryID)
	}
	if filter.ExpertiseName != "" {
		query = query.Joins("JOIN service_topics t ON t.service_id = s.id")
		query = query.Joins("JOIN expertises e ON e.id = t.expertise_id ")
		query = query.Where("e.name ilike ? ", "%"+filter.ExpertiseName+"%")
	}
	query = query.Preload("Services").Preload("Services.Category").Preload("Services.SubCategory").Preload("User").Preload("Image").Offset(filter.Offset).Limit(filter.Limit).Find(&talents)
	return talents
}

// AddService used to create new service for specific talent
func AddService(service *e.Service) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&service).Error; err != nil {
			return err
		}
		return nil
	})
}
