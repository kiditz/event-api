package router

import (
	c "github.com/kiditz/spgku-job/controller"
	"github.com/labstack/echo/v4"
)

// InitRoutes Initialize Routing
func InitRoutes(e *echo.Echo) {
	e.POST("/event", c.NewEvent)
}
