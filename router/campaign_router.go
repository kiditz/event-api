package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetCampaignRoutes to initialize routing used by event
func SetCampaignRoutes(v1 *echo.Group) {
	v1.POST("/campaigns", c.AddCampaign, m.IsLoggedIn(), m.IsCompany)
	v1.POST("/campaigns/documents", c.AddDocument, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/campaigns/:id", c.FindcampaignByID, m.IsLoggedIn())
	v1.GET("/campaigns/date", c.GetCampaignByDate, m.IsLoggedIn())
	v1.GET("/campaigns/social-media", c.GetAllSocialMedia, m.IsLoggedIn())
}
