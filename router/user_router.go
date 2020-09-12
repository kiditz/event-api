package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetUserRoutes to initialize routing used by user
func SetUserRoutes(v1 *echo.Group) {
	v1.POST("/user", c.AddUser)
	v1.GET("/account", c.FindUserByLoggedIn, m.IsLoggedIn())
	v1.GET("/account/:id", c.FindUserByID)
	v1.PUT("/user", c.EditUser, m.IsLoggedIn())
	v1.GET("/users", c.GetUsers)
	v1.POST("/auth/token", c.SignIn)
	v1.GET("/user/private", c.TestClaims, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/user/incomes", c.GetIncomes, m.IsLoggedIn(), m.IsTalent)
	v1.GET("/user/incomes/total", c.FindIncomeInfo, m.IsLoggedIn(), m.IsTalent)
	v1.GET("/banks", c.GetBanks)
	v1.POST("/user/bank", c.AddUserBank, m.IsLoggedIn(), m.IsTalent)
	v1.GET("/user/banks", c.GetUserBanks, m.IsLoggedIn(), m.IsTalent)

}
