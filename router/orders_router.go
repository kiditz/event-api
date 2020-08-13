package router

import (
	c "github.com/kiditz/spgku-api/controller"
	"github.com/labstack/echo/v4"
)

// SetOrderRoutes to initialize routing cart
func SetOrderRoutes(v1 *echo.Group) {
	v1.POST("/cart", c.AddCart)

}
