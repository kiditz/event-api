package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetUserRoutes to initialize routing used by user
func SetUserRoutes(v1 *echo.Group) {
	v1.POST("/user", c.AddUser)
	v1.POST("/auth/token", c.SignIn)
	v1.GET("/user/private", c.TestClaims, m.IsLoggedIn(), m.IsCompany)

}
