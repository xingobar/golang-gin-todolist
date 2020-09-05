package util

import (
	"github.com/jinzhu/gorm"
	"golang-gin-todolist/pkg/setting"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Paginator struct {
	Total int `json:"total"` // 總筆數
	TotalPage int `json:"total_page"` // 總頁數
	Data interface{} `json:"data"` // 資料
	Page int `json:"page"` // 目前頁碼
}

// 分頁 scope
func Paginate(page int) func(db *gorm.DB) *gorm.DB {
	perPage := setting.TWENTY_PAGE
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - setting.TWENTY_PAGE) * perPage).Limit(perPage)
	}
}

// 創建分頁
func CreatePaginate(total int, data interface{}, page int) *Paginator {
	return &Paginator{
		Total: total,
		Data: data,
		TotalPage: total / setting.TWENTY_PAGE,
		Page: page,
	}
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