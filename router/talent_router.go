package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetTalentRoutes to initialize routing talent
func SetTalentRoutes(v1 *echo.Group) {
	v1.POST("/talents", c.AddTalent, m.IsLoggedIn(), m.IsTalent)
	v1.GET("/campaigns/:id", c.FindCampaignByID, m.IsLoggedIn())
}
