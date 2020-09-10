package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// AddUser godoc
func AddUser(user *e.User) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		user.Active = true
		if user.Type == "talent" {
			tx.Save(&e.Talent{
				UserID:     user.ID,
				IsVerified: false,
			})
		}
		if user.Type == "company" {
			tx.Save(&e.Company{
				UserID:      user.ID,
				Name:        "",
				Description: "[{\"insert\":\"...\"},{\"insert\":\"\n\"}]",
				IsUpdated:   false,
			})
		}
		return nil
	})
}

// GetUsersByService docs
func GetUsersByService(c echo.Context, filter *e.FilteredUsers) []e.User {
	if filter.Limit == 0 {
		filter.Limit = 20
	}

	var users []e.User
	query := db.DB
	query = query.Select("DISTINCT users.*")
	query = query.Joins("JOIN services ON users.id = services.user_id AND users.active = ?", true)
	query = query.Preload("Services", func(db *gorm.DB) *gorm.DB {
		return db.Order("services.price ASC")
	}).Preload("Services.Category").Preload("Services.SubCategory")

	query = query.Preload("Talent.Location")

	if filter.Query != "" {
		query = query.Where("users.name ilike ? or users.email like ? ", "%"+filter.Query+"%", "%"+filter.Query+"%")
	}
	if filter.SubCategoryID > 0 {
		query = query.Where("services.sub_category_id = ?", filter.SubCategoryID)
	}
	if filter.CategoryID > 1 {
		query = query.Where("services.category_id = ?", filter.CategoryID)
	}
	if filter.BriefID > 0 {
		query = query.Joins("JOIN briefs b ON b.id = ?", filter.BriefID)
		query = query.Where("NOT EXISTS (SELECT u.name FROM users u JOIN services s ON u.id = s.user_id JOIN invitations i ON i.service_id = s.id WHERE u.id = users.id AND i.brief_id =  b.id)")
		query = query.Where(`NOT EXISTS (SELECT u.name FROM users u JOIN services s ON u.id = s.user_id JOIN quotations q ON q.service_id = s.id WHERE u.id = users.id AND q.brief_id =  b.id)`)
	}

	if filter.ExpertiseName != "" {
		query = query.Joins("JOIN service_topics ON services.id = service_topics.service_id")
		query = query.Joins("JOIN expertises ON expertises.id = service_topics.expertise_id")
		query = query.Where("expertises.name ILIKE ?", "%"+filter.ExpertiseName+"%")
	}
	query = query.Where("services.status = ?", e.APPROVED)
	query = query.Offset(filter.Offset).Limit(filter.Limit).Find(&users)
	return users
}

//EditUser godoc
func EditUser(c echo.Context, user *e.User) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		currentUser, err := FindUserByID(user.ID)
		if err != nil {
			return err
		}
		now := time.Now().UTC()
		user.Password = currentUser.Password
		user.Type = currentUser.Type
		user.UpdatedAt = &now
		user.UpdatedBy = utils.GetEmail(c)
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindUserByEmail is used to query user by email address
func FindUserByEmail(email string) (e.User, error) {
	var user e.User
	if db.DB.Where("email = ?", email).Find(&user).RecordNotFound() {
		err := fmt.Errorf("user_not_found")
		return user, err
	}
	return user, nil
}

// FindUserByID is used to query user by email address
func FindUserByID(userID uint) (e.User, error) {
	var user e.User
	query := db.DB
	query = query.Select("DISTINCT users.*")
	query = query.Joins("LEFT JOIN services ON users.id = services.user_id")
	query = query.Preload("Services", func(db *gorm.DB) *gorm.DB {
		return db.Order("services.price ASC")
	}).Preload("Services.Category").Preload("Services.SubCategory").Preload("Services.Background")

	query = query.Preload("Talent.Location").Preload("Talent.Expertises").Preload("Talent.Occupations")

	if query.Where("users.id = ?", userID).Find(&user).RecordNotFound() {
		err := fmt.Errorf("user_not_found")
		return user, err
	}
	return user, nil
}
