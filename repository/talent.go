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
