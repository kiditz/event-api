package translate

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"io/ioutil"

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
	// Handle translation from en.json
	addTranslator(transEn, "en.json")
	// Handle translation from id.json
	addTranslator(transID, "id.json")

}

func addTranslator(trans ut.Translator, fileName string) {
	file, _ := ioutil.ReadFile(fileName)
	var data map[string]string
	json.Unmarshal([]byte(file), &data)
	for key, value := range data {
		trans.Add(key, value, false)
	}
}

// Translate the value
func Translate(c echo.Context, s interface{}) ([]map[string]interface{}, ut.Translator) {
	err := VALIDATE.Struct(s)

	var slice []map[string]interface{}
	trans := GetTranslator(c)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			mapper := make(map[string]interface{})
			mapper["pascal_field"] = err.StructField()
			mapper["field"] = strings.ToLower(err.Field())
			mapper["value"] = strings.TrimSpace(strings.ReplaceAll(err.Translate(trans), err.Field(), ""))
			slice = append(slice, mapper)
		}
		return slice, trans
	}
	return slice, trans
}

// GetTranslator is used to call translator
func GetTranslator(c echo.Context) ut.Translator {
	acceptLanguage := c.Request().Header.Get("Accept-Language")
	trans, _ := UNI.GetTranslator(acceptLanguage)

	return trans
}

// TranslateError to get error from translate.json into string value
func TranslateError(c echo.Context, err error) string {
	tx := GetTranslator(c)
	res, _ := tx.T(err.Error())
	return res
}

// Errors is used for handle error by standarize api
func Errors(c echo.Context, statusCode int, s interface{}) error {
	if reflect.TypeOf(s).Name() == "string" {
		return c.JSON(statusCode, ErrorModel{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
			Message:    s.(string),
		})
	}
	return c.JSON(statusCode, ErrorModel{
		Status:     http.StatusText(statusCode),
		Errors:     s,
		StatusCode: statusCode,
	})
}

// Success is used to handle data success
func Success(c echo.Context, s interface{}) error {
	return c.JSON(http.StatusOK, ErrorModel{
		Data:       s,
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	})
}
