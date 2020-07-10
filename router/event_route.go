package router

import (
	c "github.com/kiditz/spgku-job/controller"
	"github.com/labstack/echo/v4"
)

// SetEventRoutes to initialize routing used by event
func SetEventRoutes(e *echo.Echo) {
	e.POST("/event", c.NewEvent)
}
