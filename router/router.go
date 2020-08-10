package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// InitRoutes Initialize Routing
func InitRoutes(v1 *echo.Group) {
	SetCampaignRoutes(v1)
	SetUserRoutes(v1)
	SetTalentRoutes(v1)
	v1.GET("/categories", c.GetCategoies)
}
