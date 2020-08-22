package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetCampaignRoutes to initialize routing used by event
func SetCampaignRoutes(v1 *echo.Group) {
	v1.POST("/campaigns", c.AddCampaign, m.IsLoggedIn(), m.IsCompany)
	v1.POST("/campaigns/documents", c.AddDocument, m.IsLoggedIn())
	v1.GET("/campaigns/:id", c.FindCampaignByID, m.IsLoggedIn())
	v1.GET("/campaigns", c.GetCampaigns, m.IsLoggedIn())
	v1.GET("/campaigns/payment-terms", c.GetPaymentTerms, m.IsLoggedIn())
	v1.GET("/campaigns/payment-days", c.GetPaymentDays, m.IsLoggedIn())
	v1.GET("/campaigns/info/:id", c.GetCampaignInfo)
}
