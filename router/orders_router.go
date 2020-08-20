package router

import (
	c "github.com/kiditz/spgku-api/controller"
	m "github.com/kiditz/spgku-api/middleware"
	"github.com/labstack/echo/v4"
)

// SetOrderRoutes to initialize routing cart
func SetOrderRoutes(v1 *echo.Group) {
	v1.POST("/cart", c.AddCart)
	v1.POST("/invitation", c.AddInvitation, m.IsLoggedIn(), m.IsCompany)
	v1.GET("/invitations", c.GetInvitations, m.IsLoggedIn(), m.IsTalent)
	v1.POST("/invitation/accept", c.AcceptInvitation, m.IsLoggedIn(), m.IsTalent)
	v1.POST("/invitation/reject", c.RejectInvitation, m.IsLoggedIn(), m.IsTalent)
	v1.DELETE("/cart", c.DeleteCart)
	v1.GET("/quotations", c.GetQuotations)
	v1.GET("/carts", c.GetCarts)

}
