package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ParentId uint `gorm:"column:parent_id" json:"parent_id"`
	UserId   uint `gorm:"column:user_id" json:"user_id"`
	Content string `gorm:"column:content" json:"content"`
	ArticleId uint `gorm:"column:article_id" json:"article_id"`
	Children []Comment `gorm:"foreignKey:ParentId" json:"children"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (Comment) TableName() string {
	return "comments"
}

// 取得父留言
func GetParentComment() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id is null")
	}
}

func GetChildComment() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id = 0").Preload("Children")
	}
}
