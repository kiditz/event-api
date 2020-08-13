package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
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
// @Summary DeleteCart api used to delete cart for specific ip address
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
	return t.Success(c, map[string]string{"ip_address": c.RealIP()})
}
