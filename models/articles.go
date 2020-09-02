package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string `gorm:"column:title;size:255"`
	Content string `gorm:"column:content"`
	UserId int `gorm:"column:user_id"`
	Tags []Tag `gorm:"many2many:article_tags;"`
	gorm.Model
}

func (Article) TableName() string{
	return "articles"
}

// 新增文章
func (a *Article) Create(article Article, tags []Tag) (error) {
	Db.Create(&article)
	if err := Db.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		return err
	}
	return nil
}
