package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetDigitalStaffRoute to initialize routing used by event
func SetDigitalStaffRoute(v1 *echo.Group) {
	v1.POST("/digital-staff", c.AddDigitalStaff)
}
