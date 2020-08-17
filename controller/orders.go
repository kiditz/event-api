package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddCart godoc
// @Summary AddCart api used to create new cart for specific email address
// @Description add to cart
// @Tags orders
// @MimeType
// @Produce json
// @Param talent body entity.Cart true "Add To Cart"
// @Success 200 {object} translate.ResultSuccess{data=entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /cart [post]
func AddCart(c echo.Context) error {
	var cart e.Cart
	err := c.Bind(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddToCart(&cart, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, cart)
}

// DeleteCart godoc
// @Summary DeleteCart api used to delete cart for specific device
// @Description add to cart
// @Tags orders
// @Param device_id query string true "Device ID"
// @MimeType
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /cart [delete]
func DeleteCart(c echo.Context) error {
	deviceID := c.QueryParam("device_id")
	err := r.DeleteCart(deviceID)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, map[string]string{"device_id": deviceID})
}

// GetCarts godoc
// @Summary GetCarts api used to find cart for specific device
// @Description find carts
// @Tags orders
// @Param device_id query string true "Device ID"
// @MimeType
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Cart} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /carts [get]
func GetCarts(c echo.Context) error {
	deviceID := c.QueryParam("device_id")
	carts := r.GetCarts(deviceID)

	return t.Success(c, carts)
}

// AddInvitation godoc
// @Summary AddInvitation api used to create new invitation for talent service
// @Description create new invitation
// @Tags orders
// @MimeType
// @Produce json
// @Param talent body entity.Invitation true "Invitation"
// @Success 200 {object} translate.ResultSuccess{data=entity.Invitation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitation [post]
func AddInvitation(c echo.Context) error {
	var invitations []e.Invitation
	err := c.Bind(&invitations)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddInvitation(&invitations)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, invitations)
}

// GetInvitations godoc
// @Summary GetInvitations api used to invitations by user logged in
// @Description find invitations
// @Tags orders
// @MimeType
// @Produce json
// @Param invitation query entity.LimitOffset false "LimitOffset"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.Invitation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /invitations [get]
// @Security ApiKeyAuth
func GetInvitations(c echo.Context) error {
	var limitOffset e.LimitOffset
	err := c.Bind(&limitOffset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	email := utils.GetEmail(c)
	invitations := r.GetInvitations(email, limitOffset)
	return t.Success(c, invitations)
}

// AddQuotation godoc
// @Summary AddQuotation api used to create new invitation for talent service
// @Description create new invitation
// @Tags orders
// @MimeType
// @Produce json
// @Param talent body entity.Quotation true "Quotation"
// @Success 200 {object} translate.ResultSuccess{data=entity.Quotation} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /quotation [post]
func AddQuotation(c echo.Context) error {
	var cart e.Cart
	err := c.Bind(&cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddToCart(&cart, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, cart)
}
