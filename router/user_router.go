package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetUserRoutes to initialize routing used by user
func SetUserRoutes(e *echo.Echo) {
	e.POST("/api/v1/user", c.AddUser)
	e.POST("/api/v1/auth/token", c.SignIn)
	e.GET("/api/v1/user/private", c.TestClaims, m.IsLoggedIn(), m.IsCompany)

}
