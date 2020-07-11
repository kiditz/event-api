package controller

import (
	"net/http"
	"time"

	e "github.com/kiditz/spgku-job/entity"
	r "github.com/kiditz/spgku-job/repository"
	t "github.com/kiditz/spgku-job/translate"
	"github.com/labstack/echo/v4"
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
		return t.Errors(c, translate)
	}
	job.CreatedAt = time.Now().UTC()
	err = r.CreateJob(job)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, job)
}
