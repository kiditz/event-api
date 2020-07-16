package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// ParseJwtString used to parse jwt token
func ParseJwtString(tokenString string) (map[string]interface{}, error) {
	base64Url := strings.Split(tokenString, ".")[1]
	var result map[string]interface{}
	base64Str := strings.Replace(base64Url, "-", "+", -1)
	base64Str = strings.Replace(base64Str, "_", "/", -1)
	data, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return result, err
}
