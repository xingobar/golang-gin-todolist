package models

import (
	"time"
)

type Comment struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ParentId uint `gorm:"column:parent_id" json:"parent_id"`
	UserId   uint `gorm:"column:user_id" json:"user_id"`
	Content string `gorm:"column:content" json:"content"`
	ArticleId uint `gorm:"column:article_id" json:"article_id"`
	Children []Comment `json:"children"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return "comments"
}
