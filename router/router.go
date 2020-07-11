package router

import (
	"github.com/labstack/echo/v4"
)

// InitRoutes Initialize Routing
func InitRoutes(e *echo.Echo) {
	SetCompanyRoutes(e)
	SetJobRoutes(e)
}
