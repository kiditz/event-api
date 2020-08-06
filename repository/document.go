package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
)

// AddDocument godoc
func AddDocument(document *e.Document) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&document).Error; err != nil {
			return err
		}
		return nil
	})
}
