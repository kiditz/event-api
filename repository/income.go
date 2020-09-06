package repository

import (
	"github.com/kiditz/spgku-api/db"
	e "github.com/kiditz/spgku-api/entity"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
)

// GetIncomes used to insert campaign into briefs database
func GetIncomes(c echo.Context, filter *e.IncomeFilter) []e.Income {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	incomes := []e.Income{}
	user := utils.GetUser(c)
	userID := uint(user["id"].(float64))
	// email := user["email"].(string)
	query := db.DB.Where("user_id = ?", userID)
	if filter.StartDate != "" && filter.EndDate != "" {
		query = query.Where("created_at between to_char(?, 'YYYY-MM-DD') and to_char(?, 'YYYY-MM-DD')", filter.StartDate, filter.EndDate)
	}
	if filter.CanWithdrawal {
		query = query.Where("can_withdrawal = ?", filter.CanWithdrawal)
	}
	if filter.HasWithdraw {
		query = query.Where("has_withdraw = ?", filter.HasWithdraw)
	}
	query = query.Preload("Brief")
	query = query.Offset(filter.Offset).Limit(filter.Limit).Order("id desc")
	query.Find(&incomes)
	return incomes
}
