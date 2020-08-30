package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Id int `json:"id"`
	Title string `json:"title"`
	gorm.Model
}