package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetCampaignRoutes to initialize routing used by event
func SetCampaignRoutes(v1 *echo.Group) {
	v1.POST("/campaigns", c.AddCampaign, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/campaigns/:id", c.FindcampaignByID, m.IsLoggedIn())
}
