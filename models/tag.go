package models

import (
	"time"
)

type Tag struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
	Articles []Article `gorm:"many2many:article_tags;" json:"articles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (Tag) TableName() string {
	return "tags"
}
