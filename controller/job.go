package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// CreateJob to create new job for the talent
func CreateJob(c echo.Context) error {
	job := new(e.Job)
	err := c.Bind(&job)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.ValidateTranslator(c, job)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	err = r.CreateJob(job)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
	}
	return t.Success(c, job)
}
