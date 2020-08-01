package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetCompanyRoutes to initialize routing used by company
func SetCompanyRoutes(v1 *echo.Group) {
	v1.POST("/company", c.CreateCompany)
	v1.GET("/company/:id", c.FindCompany)
	v1.GET("/company", c.GetCompanies)
}
