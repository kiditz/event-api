package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetJobRoutes to initialize routing used by event
func SetJobRoutes(v1 *echo.Group) {
	v1.POST("/campaign", c.AddCampaign, m.IsLoggedIn(), m.IsCompany)
}
