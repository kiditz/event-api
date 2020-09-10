package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddTalent godoc
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
		tx.Model(&talent).Association("Occupations").Replace(talent.Occupations)
		return nil
	})
}

// FindTalentByID  godoc
func FindTalentByID(talentID int) (e.Talent, error) {
	var talent e.Talent
	query := db.DB.Where("id=?", talentID)
	// query = query.Preload("Services")
	query = query.Preload("Expertises")
	query = query.Preload("User")
	query = query.Preload("BusinessType")
	query = query.Preload("Location")
	if query.Find(&talent).RecordNotFound() {
		return talent, fmt.Errorf("talent_not_found")
	}

	return talent, nil
}

// FindTalentByEmail godoc
func FindTalentByEmail(email string) (e.Talent, error) {
	var talent e.Talent
	query := db.DB.Joins("JOIN users u ON u.id = talents.user_id").Where("u.email = ?", email)
	query = query.Preload("User.Services").Preload("User.Services.Background")
	query = query.Preload("User.Services.Category").Preload("User.Services.SubCategory").Preload("User.Services.Topics")
	query = query.Preload("Expertises")
	query = query.Preload("Image").Preload("BackgroundImage")
	query = query.Preload("User")
	query = query.Preload("BusinessType")
	query = query.Preload("Location")
	query = query.Preload("Occupations")
	if query.First(&talent).RecordNotFound() {
		return talent, fmt.Errorf("talent_not_found")
	}
	return talent, nil
}

// GetTalentList  godoc
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
	if filter.BriefID > 0 {
		query = query.Where("NOT EXISTS (SELECT 1 FROM quotations i WHERE i.brief_id = ? AND i.service_id = s.id)", filter.BriefID)
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

// AddService godoc
func AddService(service *e.Service, c echo.Context) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		user := utils.GetUser(c)
		userID := uint(user["id"].(float64))
		service.UserID = userID
		service.Status = e.ONREVIEW
		tx.Model(&service).Association("Topics").Replace(service.Topics)
		tx.Model(&service).Association("Background").Replace(service.Background)
		tx.Model(&service).Association("Portfilios").Replace(service.Portfilios)

		if err := tx.Save(&service).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindServiceByID godoc
func FindServiceByID(serviceID int) (e.Service, error) {
	var service e.Service
	db := db.DB
	db = db.Preload("Category").Preload("SubCategory")
	db = db.Preload("User.Talent.Location")
	db = db.Preload("Topics")
	db = db.Preload("Portfilios")
	db = db.Preload("Background")

	if db.First(&service, serviceID).RecordNotFound() {
		err := fmt.Errorf("service_not_found")
		return service, err
	}
	return service, nil
}
