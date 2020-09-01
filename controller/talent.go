package controller

import (
	"net/http"
	"strconv"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddTalent godoc
// @Summary AddTalent api used to create new talent
// @Description Create a new talent
// @Tags talents
// @MimeType
// @Produce json
// @Param talent body entity.Talent true "New Talent"
// @Success 200 {object} translate.ResultSuccess{data=entity.Talent} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talent [post]
// @Security ApiKeyAuth
func AddTalent(c echo.Context) error {
	var talent e.Talent
	err := c.Bind(&talent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = r.AddTalent(&talent, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, talent)
}

// AddService godoc
// @Summary AddService api used to create new service for talent
// @Description Create a new service
// @Tags talents
// @MimeType
// @Produce json
// @Param service body entity.Service true "New Service for talent"
// @Success 200 {object} translate.ResultSuccess{data=entity.Service} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talent/service [post]
// @Security ApiKeyAuth
func AddService(c echo.Context) error {
	var service e.Service
	err := c.Bind(&service)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// TODO: Service add
	err = r.AddService(&service, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, service)
}

// FindTalentByID godoc
// @Summary FindtalentById used to find talent by it's primary key
// @Description find talent by id
// @Tags talents
// @Accept json
// @Produce json
// @Param id path string true "Talent ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.Talent} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talent/{id} [get]
// @Security ApiKeyAuth
func FindTalentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	talent, err := r.FindTalentByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, talent)
}

// FindTalentByLogin godoc
// @Summary FindTalentByLogin used to find talent login
// @Description find talent login
// @Tags talents
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=entity.Talent} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talent [get]
// @Security ApiKeyAuth
func FindTalentByLogin(c echo.Context) error {
	email := utils.GetEmail(c)
	talent, err := r.FindTalentByEmail(email)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, talent)
}

// GetTalents godoc
// @Summary GetTalents is api to find talents by params
// @Description find talents
// @Tags talents
// @Accept json
// @Produce json
// @Param filter query entity.FilteredTalent false "FilteredTalent"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.TalentResults} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talents [get]
// @Security ApiKeyAuth
func GetTalents(c echo.Context) error {
	filter := new(e.FilteredTalent)
	if err := c.Bind(filter); err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	talents := r.GetTalentList(filter)
	return t.Success(c, talents)
}

// FindServiceByID godoc
// @Summary FindServiceByID used to find service by primary key
// @Description find service by id
// @Tags talents
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.Service} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /talent/service/{id} [get]
// @Security ApiKeyAuth
func FindServiceByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	service, err := r.FindServiceByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, service)
}
