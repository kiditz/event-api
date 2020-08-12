package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddTalent used to create new talent
func AddTalent(talent *e.Talent, c echo.Context) error {

	return db.DB.Transaction(func(tx *gorm.DB) error {
		user := utils.GetUser(c)
		talent.UserID = uint(user["id"].(float64))
		talent.CreatedBy = user["email"].(string)
		tx = tx.Set("gorm:association_autoupdate", false)
		if err := tx.Save(&talent).Error; err != nil {
			return err
		}

		tx.Model(&talent.Expertises).Association("Expertises").Append(talent.Expertises)
		if talent.BackgroundImage != nil {
			tx.Model(&talent.BackgroundImage).Save(talent.BackgroundImage)
			tx.Model(&talent).Association("BackgroundImage").Append(talent.BackgroundImage)
		}
		if talent.Image != nil {
			tx.Model(&talent.Image).Save(talent.Image)
			tx.Model(&talent).Association("Image").Append(talent.Image)
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

// FindTalentByEmail used to find talent by email
func FindTalentByEmail(email string) (e.Talent, error) {
	var talent e.Talent
	if err := db.DB.Joins("JOIN users u ON u.id = talents.user_id").Where("u.email = ?", email).First(&talent).Error; err != nil {
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
func AddService(service *e.Service, c echo.Context) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		tx = tx.Set("gorm:association_autoupdate", false)
		email := utils.GetEmail(c)
		talent, err := FindTalentByEmail(email)
		if err != nil {
			return err
		}
		service.TalentID = talent.ID
		if err := tx.Save(&service).Error; err != nil {
			return err
		}
		return nil
	})
}
