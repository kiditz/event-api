package utils

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
)

//GetUser used to getting the user from claims
func GetUser(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// fmt.Printf("claims %s", claims)
	return claims
}

//GetEmail used to get user email
func GetEmail(c echo.Context) string {
	return GetUser(c)["email"].(string)
}
