package controller

import (
	"net/http"
	"strconv"

	e "github.com/kiditz/spgku-job/entity"
	r "github.com/kiditz/spgku-job/repository"
	t "github.com/kiditz/spgku-job/translate"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// CreateCompany used to insert company data into database
func CreateCompany(c echo.Context) error {
	company := new(e.Company)
	err := c.Bind(&company)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.Translate(c, company)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	err = r.CreateCompany(company)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, company)
}

// FindCompany used to found company by it's primary key id
func FindCompany(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	company, err := r.FindCompanyByID(id)
	if err != nil {
		return t.Errors(c, http.StatusNotFound, t.TranslateError(c, err))
	}
	return t.Success(c, company)
}
