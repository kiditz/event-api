package controller

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddCampaign to create new campaign
func AddCampaign(c echo.Context) error {
	campaign := new(e.Campaign)
	err := c.Bind(&campaign)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.ValidateTranslator(c, campaign)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Printf("claims %s", claims)
	campaign.CreatedBy = claims["email"].(string)
	err = r.AddCampaign(campaign)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
	}
	return t.Success(c, campaign)
}
