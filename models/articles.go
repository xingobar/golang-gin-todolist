package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"`
	Title string `gorm:"column:title;size:255"`
	Content string `gorm:"column:content"`
	UserId int `gorm:"column:user_id"`
	Tags []Tag `gorm:"many2many:article_tags"`
	gorm.Model
}

func (Article) TableName() string{
	return "articles"
}

// 新增文章
func (a *Article) Create(article Article, tags []Tag) (error) {
	tx := Db.Begin()
	tx.Create(&article)
	if err := tx.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 取得單一文章
func (a *Article) GetById(id string) *Article {
	var article Article
	Db.Where("id = ?", id).First(&article)

	var tag []Tag
	if err := Db.Model(&article).Association("Tags").Find(&tag).Error; err != nil {
		return nil
	}
	article.Tags = tag
	return &article
}
