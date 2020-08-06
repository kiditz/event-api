package controller

import (
	"net/http"
	"os"

	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/labstack/echo/v4"
)

// AddDocument godoc
// @Summary Upload documents
// @Description Upload file
// @Tags campaigns
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "campaign image"
// @Success 200 {object} translate.ResultSuccess{data=entity.Campaign} desc
// @Failure 400 {object} translate.ResultErrors
// @Failure 404 {object} translate.ResultErrors
// @Failure 500 {object} translate.ResultErrors
// @Router /campaigns/documents [post]
// @Security ApiKeyAuth
func AddDocument(c echo.Context) error {
	// Get avatar
	file, err := c.FormFile("file")
	if err != nil {
		return t.Errors(c, http.StatusInternalServerError, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return t.Errors(c, http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	document := &e.Document{
		FileName: file.Filename,
		Size:     file.Size,
	}
	err = r.AddDocument(document)
	if err == nil {
		return t.Success(c, document)
	}
	return t.Errors(c, http.StatusInternalServerError, map[string]string{"error": http.StatusText(http.StatusInternalServerError)})
}
