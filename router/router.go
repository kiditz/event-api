package router

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes Initialize Routing
func InitRoutes(v1 *echo.Group) {
	SetCompanyRoutes(v1)
	SetJobRoutes(v1)
	SetUserRoutes(v1)
	SetDigitalStaffRoute(v1)
}
