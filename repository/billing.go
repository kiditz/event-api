package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// GetBillings godoc
func GetBillings(c echo.Context, filter *e.BillingFilter) []e.Billing {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	billings := []e.Billing{}
	user := utils.GetUser(c)
	userID := uint(user["id"].(float64))

	query := db.DB.Where("user_id = ?", userID)
	if filter.StartDate != "" && filter.EndDate != "" {
		query = query.Where("due_date between to_char(?, 'YYYY-MM-DD') and to_char(?, 'YYYY-MM-DD')", filter.StartDate, filter.EndDate)
	}

	if filter.Query != "" {
		query = query.Joins("JOIN briefs b ON b.id = brief_id")
		query = query.Where("order_id ilike ? OR b.title ilike ?", "%"+filter.Query+"%")
	}
	query = query.Preload("Brief")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc")
	query.Find(&billings)
	return billings
}
