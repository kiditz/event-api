package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetStaffRoute to initialize routing used by event
func SetStaffRoute(v1 *echo.Group) {
	v1.POST("/digital-staffs", c.AddDigitalStaff)
	v1.POST("/event-staffs", c.AddEventStaff)
	v1.GET("/digital-staffs", c.GetDigitalStaff)
	v1.GET("/event-staffs", c.GetEventStaff)
}
