package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	u "github.com/kiditz/spgku-api/utils"
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
	user.CreatedBy = u.GetUsername(c)
	pwd, err := u.HashAndSalt([]byte(user.Password))
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	user.Password = pwd
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

//SignIn used to login
func SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := r.FindUserByEmail(email)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	fmt.Printf("Name : %v", user.Name)
	if u.ComparePasswords(user.Password, []byte(password)) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["email"] = user.Email
		claims["type"] = user.Type
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		result, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			return err
		}
		return t.Success(c, map[string]string{"token": result})
	}
	return t.Errors(c, http.StatusUnauthorized, t.TranslateString(c, "user_not_found"))
}

// TesTestClaims test private
func TestClaims(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	return t.Success(c, claims)
}
