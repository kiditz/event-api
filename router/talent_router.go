package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetTalentRoutes to initialize routing talent
func SetTalentRoutes(v1 *echo.Group) {
	v1.POST("/talent", c.AddTalent, m.IsLoggedIn(), m.IsTalent)
	v1.POST("/talent/service", c.AddService, m.IsLoggedIn(), m.IsTalent)
	v1.GET("/talents", c.GetTalents)
	v1.GET("/talent/:id", c.FindTalentByID)
	v1.GET("/talent", c.FindTalentByLogin, m.IsLoggedIn(), m.IsTalent)
}
