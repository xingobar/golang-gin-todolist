package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(data []byte) string {
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedData string, plaintext []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedData)
	err := bcrypt.CompareHashAndPassword(byteHash, plaintext)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}