package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	e "github.com/kiditz/spgku-api/entity"
	r "github.com/kiditz/spgku-api/repository"
	t "github.com/kiditz/spgku-api/translate"
	"github.com/kiditz/spgku-api/utils"
	u "github.com/kiditz/spgku-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

// AddUser godoc
// @Summary AddUser api used to signup
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.UserForm true "New User"
// @Success 200 {object} translate.ResultSuccess{data=entity.User} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /user [post]
func AddUser(c echo.Context) error {
	form := new(e.UserForm)
	err := c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	pwd, err := u.HashAndSalt([]byte(form.Password))
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	form.Password = pwd
	userData := e.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
		Type:     form.Type,
	}
	err = r.AddUser(&userData)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Error())
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, userData)
}

// EditUser godoc
// @Summary EditUser api used to edit profile
// @Description Edit user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entity.User true "Edit User"
// @Success 200 {object} translate.ResultSuccess{data=entity.User} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /user [put]
// @Security ApiKeyAuth
func EditUser(c echo.Context) error {
	user := new(e.User)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = r.EditUser(c, user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			return t.Errors(c, http.StatusBadRequest, err.Error())
		}
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, user)
}

// FindUserByLoggedIn godoc
// @Summary FindUserByLoggedIn api used to find user by token
// @Description Find user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} translate.ResultSuccess{data=entity.User} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /account [get]
// @Security ApiKeyAuth
func FindUserByLoggedIn(c echo.Context) error {
	email := utils.GetEmail(c)
	user, err := r.FindUserByEmail(email)
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, user)
}

// FindUserByID godoc
// @Summary FindUserByID api used to find user by primary key
// @Description Find user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} translate.ResultSuccess{data=entity.User} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /account/{id} [get]
// @Security ApiKeyAuth
func FindUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := r.FindUserByID(uint(id))
	if err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	return t.Success(c, user)
}

//SignIn used to login
// @Summary Sign In
// @Description Sign in by using email and password
// @Tags users
// @Accept x-www-form-urlencoded
// @Produce  json
// @Param email query string true "your account email"
// @Param password query string true "your account password"
// @Success 200 {object} translate.ResultSuccess{data=entity.Brief} desc
// @Failure 400 {object} translate.ResultErrors
// @Failure 404 {object} translate.ResultErrors
// @Failure 500 {object} translate.ResultErrors
// @Router /auth/token [post]
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
		claims["id"] = user.ID
		claims["email"] = user.Email
		claims["type"] = user.Type
		claims["currency"] = user.Currency
		claims["language"] = user.Language
		claims["image_url"] = user.ImageURL
		claims["background_image_url"] = user.BackgroundImageURL
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		result, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			return err
		}
		return t.Success(c, map[string]string{"token": result})
	}
	return t.Errors(c, http.StatusUnauthorized, t.TranslatesString(c, "user_not_found"))
}

// TestClaims test private
func TestClaims(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// name := claims["name"].(string)
	return t.Success(c, claims)
}

// GetUsers godoc
// @Summary GetUsers is api to find users by service params
// @Description find users
// @Tags talents
// @Accept json
// @Produce json
// @Param filter query entity.FilteredUsers false "FilteredUsers"
// @Success 200 {object} translate.ResultSuccess{data=[]entity.User} desc
// @Failure 400 {object} translate.ResultErrors
// @Router /users [get]
// @Security ApiKeyAuth
func GetUsers(c echo.Context) error {
	filter := new(e.FilteredUsers)
	if err := c.Bind(filter); err != nil {
		return t.Errors(c, http.StatusBadRequest, err.Error())
	}
	talents := r.GetUsersByService(filter)
	return t.Success(c, talents)
}
