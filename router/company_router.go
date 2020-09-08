package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetCompanyRoutes to initialize routing used by company
func SetCompanyRoutes(v1 *echo.Group) {
	v1.PUT("/company", c.UpdateCompany, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/company", c.FindCompany, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/company/billings", c.GetBilling, m.IsLoggedIn(), m.IsCompany)
}
