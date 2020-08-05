package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddCampaign godoc
// @Summary AddCampaign api used to signup
// @Description Create a new user
// @Tags campaigns
// @Accept json
// @Produce json
// @Param campaign body entity.Campaign true "New Campaign"
// @Success 200 {object} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns [post]
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

// FindcampaignByID godoc
// @Summary FindcampaignById used to find campaign by it's primary key
// @Description Get all data of event staff
// @Tags staff
// @Accept json
// @Produce json
// @Success 200 {array} entity.Campaign desc
// @Failure 400 {object} translate.ResultErrors
// @Router /campaigns/ [get]// GetEventStaff used to get all digital staff
// @Summary GetEventStaff used to categories help event staff
// @Description Get all data of event staff
// @Tags staff
// @Accept json
// @Param id path int true "Account ID"
// @Success 200 {array} entity.EventStaff desc
// @Failure 400 {object} translate.ResultErrors
// @Router /event-staffs [get]
func FindcampaignByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	campaign, err := r.FindCampaignByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, campaign)
}
