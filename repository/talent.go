package repository

import (
	"fmt"

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
		// email := user["email"].(string)
		userID := uint(user["id"].(float64))

		account, _ := FindUserByID(userID)
		talent.UserID = account.ID
		talent.CreatedBy = account.Email
		talent.User.Password = account.Password
		// Only need for update event
		talent.User.UpdatedBy = account.Email
		talent.User.CreatedBy = account.Email

		if err := tx.Save(&talent).Error; err != nil {
			return err
		}
		tx.Model(&talent).Association("Expertises").Replace(talent.Expertises)
		return nil
	})
}

// FindTalentByID  used to find talent by id
func FindTalentByID(talentID int) (e.Talent, error) {
	var talent e.Talent
	query := db.DB.Where("id=?", talentID)
	query = query.Preload("Services").Preload("Services.Portofilios").Preload("Services.Category").Preload("Services.SubCategory")
	query = query.Preload("Expertises")
	query = query.Preload("Image").Preload("BackgroundImage")
	query = query.Preload("User")
	query = query.Preload("BusinessType")
	query = query.Preload("Location")
	if err := query.Find(&talent).Error; err != nil {
		return talent, err
	}

	return talent, nil
}

// FindTalentByEmail used to find talent by email
func FindTalentByEmail(email string) (e.Talent, error) {
	var talent e.Talent
	query := db.DB.Joins("JOIN users u ON u.id = talents.user_id").Where("u.email = ?", email)
	query = query.Preload("Services").Preload("Services.Portofilios").Preload("Services.Category").Preload("Services.SubCategory")
	query = query.Preload("Expertises")
	query = query.Preload("Image").Preload("BackgroundImage")
	query = query.Preload("User")
	query = query.Preload("BusinessType")
	query = query.Preload("Location")
	if query.First(&talent).RecordNotFound() {
		return talent, fmt.Errorf("talent_not_found")
	}
	return talent, nil
}

// FilteredTalent used to filter talent query

// GetTalents  used to find talent by id
func GetTalents(filter *e.FilteredTalent) []e.Talent {
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
	if filter.Q != "" {
		query = query.Joins("JOIN users u ON u.id = talents.user_id ")
		query = query.Where("u.name ilike ? or u.email like ? ", "%"+filter.Q+"%", "%"+filter.Q+"%")
	}
	query = query.Preload("Services").Preload("Services.Category").Preload("Services.SubCategory").Preload("User").Preload("Image").Preload("BackgroundImage").Offset(filter.Offset).Limit(filter.Limit).Find(&talents)
	return talents
}

// GetTalentList  used to find talent by id
func GetTalentList(filter *e.FilteredTalent) []e.TalentResults {
	if filter.Limit == 0 {
		filter.Limit = 20
	}
	query := db.DB.Table("users u")
	query = query.Select("DISTINCT s.start_price, s.id, t.id, u.name , c.name, sc.name, s.image_url")
	query = query.Joins("JOIN talents t ON u.id = t.user_id ")
	query = query.Joins("LEFT JOIN talent_expertises te ON t.id = te.talent_id ")
	query = query.Joins("JOIN services s ON t.id = s.talent_id")
	query = query.Joins("JOIN categories c ON c.id = s.category_id")
	query = query.Joins("JOIN sub_categories sc ON sc.id = s.sub_category_id")
	query = query.Joins("LEFT JOIN expertises e ON e.id = te.expertise_id ")
	if filter.Q != "" {
		query = query.Where("u.name ilike ? or u.email like ? ", "%"+filter.Q+"%", "%"+filter.Q+"%")
	}
	if filter.ExpertiseName != "" {
		query = query.Where("e.name ilike ? ", "%"+filter.ExpertiseName+"%")
	}
	if filter.SubCategoryID > 0 {
		query = query.Where("s.sub_category_id = ?", filter.SubCategoryID)
	}
	if filter.CategoryID > 1 {
		query = query.Where("s.category_id = ?", filter.CategoryID)
	} else {
		query = query.Where("s.id = (SELECT max(id) FROM services WHERE talent_id = t.id)")
	}
	if filter.CampaignID > 0 {
		query = query.Where("NOT EXISTS (SELECT 1 FROM quotations i WHERE i.campaign_id = ? AND i.service_id = s.id)", filter.CampaignID)
	}
	rows, err := query.Offset(filter.Offset).Limit(filter.Limit).Order("t.id desc").Rows()
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	results := []e.TalentResults{}
	for rows.Next() {
		result := e.TalentResults{}
		err := rows.Scan(&result.StartPrice, &result.ServiceID, &result.TalentID, &result.Name, &result.CategoryName, &result.SubCategoryName, &result.ImageURL)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, result)
	}

	return results
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
