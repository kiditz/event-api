package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetBriefRoutes to initialize routing used by event
func SetBriefRoutes(v1 *echo.Group) {
	v1.POST("/briefs", c.AddBrief, m.IsLoggedIn(), m.IsCompany)
	v1.POST("/briefs/documents", c.AddDocument, m.IsLoggedIn())
	v1.GET("/briefs/:id", c.FindBriefByID, m.IsLoggedIn())
	v1.GET("/briefs", c.GetBriefs, m.IsLoggedIn())
	v1.GET("/briefs/payment-terms", c.GetPaymentTerms, m.IsLoggedIn())
	v1.GET("/briefs/payment-days", c.GetPaymentDays, m.IsLoggedIn())
	v1.GET("/briefs/info/:id", c.GetBriefInfo, m.IsLoggedIn())
}
