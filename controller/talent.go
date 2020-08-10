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
// @Router /talents [post]
// @Security ApiKeyAuth
func AddTalent(c echo.Context) error {
	talent := new(e.Talent)
	err := c.Bind(&talent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user := utils.GetUser(c)
	talent.UserID = user["id"].(uint)
	talent.CreatedBy = user["email"].(string)
	err = r.AddTalent(talent)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Constraint)
		}
	}
	return t.Success(c, talent)
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
// @Router /talents/{id} [get]
// @Security ApiKeyAuth
func FindTalentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	talent, err := r.FindTalentByID(id)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, talent)
}
