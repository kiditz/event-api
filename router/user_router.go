package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetUserRoutes to initialize routing used by user
func SetUserRoutes(e *echo.Echo) {
	e.POST("/api/v1/user", c.AddUser)

}
