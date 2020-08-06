package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kiditz/spgku-api/translate"

	"github.com/labstack/echo/v4"
)

// ParseJwtString used to parse jwt token
func ParseJwtString(tokenString string) (map[string]interface{}, error) {
	if translate.Empty(tokenString) {
		return nil, nil
	}
	base64Url := strings.Split(tokenString, ".")[1]
	var result map[string]interface{}
	base64Str := strings.ReplaceAll(base64Url, "-", "+")
	base64Str = strings.ReplaceAll(base64Str, "_", "/")
	data, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}

// ParseJwt is used for parsing JWT from header
func ParseJwt(c echo.Context) (map[string]interface{}, error) {
	auth := c.Request().Header.Get("Authorization")
	auth = strings.TrimSpace(strings.ReplaceAll(auth, "Bearer", ""))
	return ParseJwtString(auth)
}

// GetUsername get the username from jwt result
func GetUsername(c echo.Context) string {
	data, _ := ParseJwt(c)
	if data["username"] != nil {
		return fmt.Sprintf("%v", data["email"])
	}
	return "System"
}

// GetUser used to getting the user from claims
func GetUser(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// fmt.Printf("claims %s", claims)
	return claims
}

// VerifyToken used to verify is token valid
func VerifyToken(c echo.Context) (*jwt.Token, error) {
	auth := c.Request().Header.Get("Authorization")
	var tokenString = strings.TrimSpace(strings.ReplaceAll(auth, "Bearer", ""))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func main() {
// 	res, err := ParseJwtString("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiJraWRpdHpiYXN0YXJhQGdtYWlsLmNvbSIsInNjb3BlIjpbIndyaXRlIiwicmVhZCJdLCJleHAiOjE1OTQ4MzczMTUsImF1dGhvcml0aWVzIjpbIlJPTEVfVFdPX0ZBQ1RPUl9BVVRIRU5USUNBVElPTl9FTkFCTEVEIiwiRU1QTE9ZRVJTIl0sImp0aSI6IjA3ZmRmNjBmLThlNzAtNDljZS05NzE0LWRmZTJiNzdlODEyYyIsImVtYWlsIjoia2lkaXR6YmFzdGFyYUBnbWFpbC5jb20iLCJjbGllbnRfaWQiOiJzcGdrdS5jb20ifQ.maWMw1I7nHj3r6QnUDV9IDeUyDvc3Q8lq8dTBH0A--k6RokY2HSwtuaSc0g-iZScPM2o2bE1wSgGbuitjKHqqjJBUDOPCzWLF5D8xwVYYb1a9_nd_pBNi-wsqsckxbd5ye6oLheOb3P2VwzB56jPbhcu0XEDzFmCt5MC0kZmJ7ipY9IEiYPwz55e0fbGGCyXfNb9p2xoihAWuI1n9fMlst4evFFRvJzyhpRkEhs0chR4QZqx0VSvDUC12yJwo32XDOHwV5TIukF3OVyu0VjvUHo5SWAYgGtkPCke6Qz0Imv1Sf3KifxYzf9UmiJRXZoxHigRgM_sTq0U0orxFwQUMQ")
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(res["user_name"])
// }
