package models

import "github.com/jinzhu/gorm"

type Tag struct {
	Id int `gorm:"primaryKey";"column:id;autoIncrement"`
	Title string `gorm: "column:title";"size:255" `
	gorm.Model
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) Add(title string) error{
	return Db.Create(&Tag{Title:title}).Error
}
