package controller

import (
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
)

// GetCategoies godoc
// @Summary GetCategoies used to find all categories
// @Description find category by date
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} translate.ResultSuccess{data=entity.Category} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /categories [get]
// @Security ApiKeyAuth
func GetCategoies(c echo.Context) error {
	categories := r.GetCategories()
	return t.Success(c, categories)
} // GetCampaigns godoc
