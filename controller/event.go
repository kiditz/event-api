package controller

import (
	"net/http"
	"time"

	e "github.com/kiditz/spgku-job/entity"
	"github.com/kiditz/spgku-job/repo"
	t "github.com/kiditz/spgku-job/trans"
	"github.com/labstack/echo/v4"
)

// NewEvent to create new event
func NewEvent(c echo.Context) error {
	event := new(e.Event)
	err := c.Bind(&event)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var tr, err = t.Translate(c, event)
	if err != nil {
		return t.ErrorHandler(c, tr)
	}
	event.CreatedAt = time.Now().UTC()
	err = repo.CreateEvent(event)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, event)
}
