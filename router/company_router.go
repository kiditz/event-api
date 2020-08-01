package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetCompanyRoutes to initialize routing used by company
func SetCompanyRoutes(e *echo.Echo) {
	e.POST("/company", c.CreateCompany)
	e.GET("/company/:id", c.FindCompany)
	e.GET("/company", c.GetCompanies)
}
