package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

// UpdateCompany godoc
// @Summary Edit company
// @Description Edit company
// @Tags company
// @MimeType
// @Produce json
// @Param company body entity.Company true "Company Data"
// @Success 200 {object} translate.ResultSuccess{data=entity.Company} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /company [put]
// @Security ApiKeyAuth
func UpdateCompany(c echo.Context) error {
	var company e.Company
	err := c.Bind(&company)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// TODO: Service add
	err = r.UpdateCompany(&company, c)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, company)
}

// FindCompany godoc
// @Summary FindCompany used to find company logged in
// @Description find company logged in
// @Tags company
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=entity.Company} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /company [get]
// @Security ApiKeyAuth
func FindCompany(c echo.Context) error {

	company, err := r.FindCompany(c)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, company)
}
