package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// GetBilling godoc
func GetBilling(c echo.Context, filter *e.IncomeFilter) []e.Billing {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	billings := []e.Billing{}
	email := utils.GetEmail(c)
	query := db.DB.Joins("JOIN users u ON u.id = billings.user_id AND u.email = ?", email)

	query = query.Preload("Brief")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc")
	query.Find(&billings)
	return billings
}
