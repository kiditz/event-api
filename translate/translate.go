package translate

import (
	"encoding/json"
	"fmt"
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

// ValidateTranslator the value
func ValidateTranslator(c echo.Context, s interface{}) []map[string]interface{} {
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

		return slice
	}
	return slice
}

// GetTranslator is used to call translator
func GetTranslator(c echo.Context) ut.Translator {
	acceptLanguage := c.Request().Header.Get("Accept-Language")
	trans, _ := UNI.GetTranslator(acceptLanguage)

	return trans
}

// Translation handle error message
func Translation(c echo.Context, err error) string {
	return TranslatesString(c, err.Error())
}

//TranslatesString is used to handle error by using input string
func TranslatesString(c echo.Context, input string) string {
	tx := GetTranslator(c)
	res, _ := tx.T(input)
	return res
}

// Errors is used for handle error by standarize api
func Errors(c echo.Context, statusCode int, s interface{}) error {
	fmt.Println(s)
	if reflect.TypeOf(s).Name() == "string" {
		return c.JSON(statusCode, ResultErrors{
			Status:     http.StatusText(statusCode),
			StatusCode: statusCode,
			Message:    TranslatesString(c, s.(string)),
		})
	}
	return c.JSON(statusCode, ResultErrors{
		Status:     http.StatusText(statusCode),
		Errors:     s,
		Message:    TranslatesString(c, "validation_errors"),
		StatusCode: statusCode,
	})
}

// Success is used to handle data success
func Success(c echo.Context, s interface{}) error {
	return c.JSON(http.StatusOK, ResultSuccess{
		Data:       s,
		Status:     http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
	})
}
