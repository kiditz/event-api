package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt is user to generate password
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return string(pwd), err
	}
	return string(hash), nil
}

// ComparePasswords compare password is verified
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
