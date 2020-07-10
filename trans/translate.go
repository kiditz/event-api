package trans

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	id_trans "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
)

var (
	// UNI handle
	UNI *ut.UniversalTranslator
	// VALIDATE handle
	VALIDATE *validator.Validate
)

// InitTranslate initialize translation
func InitTranslate(e *echo.Echo) {
	VALIDATE = validator.New()
	VALIDATE.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
	en := en.New()
	UNI = ut.New(en, en, id.New())
	transEn, _ := UNI.GetTranslator("en")
	transID, _ := UNI.GetTranslator("id")
	en_trans.RegisterDefaultTranslations(VALIDATE, transEn)
	id_trans.RegisterDefaultTranslations(VALIDATE, transID)
}

// Translate the value
func Translate(c echo.Context, s interface{}) ([]map[string]interface{}, error) {
	err := VALIDATE.Struct(s)

	var slice []map[string]interface{}
	if err != nil {
		acceptLanguage := c.Request().Header.Get("Accept-Language")
		trans, _ := UNI.FindTranslator(acceptLanguage)
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			mapper := make(map[string]interface{})
			mapper["pascal_field"] = err.StructField()
			mapper["field"] = strings.ToLower(err.Field())
			mapper["value"] = strings.TrimSpace(strings.ReplaceAll(err.Translate(trans), err.Field(), ""))
			slice = append(slice, mapper)
		}
		return slice, err
	}
	return slice, nil
}

//Errors is used for handle error by standarize api
func Errors(c echo.Context, s interface{}) error {
	return c.JSON(http.StatusBadRequest, ErrorModel{
		Status:     http.StatusText(http.StatusBadRequest),
		Errors:     s,
		StatusCode: http.StatusBadRequest,
	})
}
