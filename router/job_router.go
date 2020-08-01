package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetJobRoutes to initialize routing used by event
func SetJobRoutes(v1 *echo.Group) {
	v1.POST("/job", c.CreateJob)
}
