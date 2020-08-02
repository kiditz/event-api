package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	u "github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddDigitalStaff godoc
// @Summary AddDigitalStaff used to categories help digital staff
// @Description Create a new digital staff category
// @Tags staff
// @Accept json
// @Produce json
// @Param digitallStaff body entity.DigitalStaff true "New DigitalStaff"
// @Success 200 {object} translate.ResultSuccess{data=entity.DigitalStaff} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /digital-staff [post]
func AddDigitalStaff(c echo.Context) error {
	digitalStaff := new(e.DigitalStaff)
	err := c.Bind(&digitalStaff)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.ValidateTranslator(c, digitalStaff)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	digitalStaff.CreatedBy = u.GetUsername(c)

	err = r.AddDigitalStaff(digitalStaff)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, digitalStaff)
}

// AddEventStaff godoc
// @Summary AddDigitalStaff used to categories help digital staff
// @Description Create a new digital staff category
// @Tags staff
// @Accept json
// @Produce json
// @Param eventStaff body entity.EventStaff true "New Event Staff"
// @Success 200 {object} translate.ResultSuccess{data=entity.EventStaff} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /event-staff [post]
func AddEventStaff(c echo.Context) error {
	eventStaff := new(e.EventStaff)
	err := c.Bind(&eventStaff)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.ValidateTranslator(c, eventStaff)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	eventStaff.CreatedBy = u.GetUsername(c)

	err = r.AddEventStaff(eventStaff)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, eventStaff)
}

//GetDigitalStaff used to get all digital staff
func GetDigitalStaff(c echo.Context) error {
	records, err := r.GetDigitalStaff()
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, records)
}

//GetEventStaff used to get all digital staff
func GetEventStaff(c echo.Context) error {
	records, err := r.GetEventStaff()
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, records)
}
