package controller

import (
	"net/http"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddUser api used to signup
func AddUser(c echo.Context) error {
	user := new(e.User)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var hasErr, tx = t.Translate(c, user)
	if hasErr != nil {
		return t.Errors(c, http.StatusBadRequest, hasErr)
	}
	user.CreatedBy = utils.GetUsername(c)
	err = r.AddUser(user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			res, _ := tx.T(err.Constraint)
			return t.Errors(c, http.StatusBadRequest, res)
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, user)
}
