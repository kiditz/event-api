package controller

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
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
	bucket := os.Getenv("BUCKET_NAME")
	object := utils.GetEmail(c) + "/" + file.Filename
	message := uploadFile(bucket, object, src)

	document := &e.Document{
		FileName: file.Filename,
		Size:     file.Size,
		URL:      "storage.googleapis.com/" + bucket + "/" + object,
		Message:  message.Error(),
	}
	err = r.AddDocument(document)
	if err == nil {
		return t.Success(c, document)
	}
	return t.Errors(c, http.StatusInternalServerError, map[string]string{"error": http.StatusText(http.StatusInternalServerError)})
}
func uploadFile(bucket string, object string, src multipart.File) error {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, src); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	acl := client.Bucket(bucket).Object(object).ACL()
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}
	return fmt.Errorf("Blob %s uploaded", object)
}
