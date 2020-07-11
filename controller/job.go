package controller

import (
	"net/http"
	"strings"

	e "github.com/kiditz/spgku-job/entity"
	r "github.com/kiditz/spgku-job/repository"
	t "github.com/kiditz/spgku-job/translate"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

const (
	// CompanyForeignError is for check if companyy foreign key error
	CompanyForeignError = "jobs_company_id_companies_id_foreign"
)

// CreateJob to create new job for the talent
func CreateJob(c echo.Context) error {
	job := new(e.Job)
	err := c.Bind(&job)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var translate = t.Translate(c, job)
	if translate != nil {
		return t.Errors(c, http.StatusBadRequest, translate)
	}
	err = r.CreateJob(job)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if strings.EqualFold(err.Constraint, CompanyForeignError) {
				return t.Errors(c, http.StatusBadRequest, err.Code.Name())
			}
			return t.Errors(c, http.StatusBadRequest, err.Error())
		}
	}
	return c.JSON(http.StatusOK, job)
}
