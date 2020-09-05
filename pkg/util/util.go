package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Paginator struct {
	Total int `json:"total"` // 總筆數
	TotalPage int `json:"total_page"` // 總頁數
	Data interface{} `json:"data"` // 資料
	Page int `json:"page"` // 目前頁碼
}

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