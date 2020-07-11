package router

import (
	c "github.com/kiditz/spgku-job/controller"
	"github.com/labstack/echo/v4"
)

// SetJobRoutes to initialize routing used by event
func SetJobRoutes(e *echo.Echo) {
	e.POST("/job", c.CreateJob)
}
