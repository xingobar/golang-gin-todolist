package models

import (
	"fmt"
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
func (a *Article) Create(article Article, tags []string) (error) {
	Db.Create(&article)
	fmt.Println("id: ", article.ID)
	tag := &Tag{}
	t ,_ := tag.GetById(1)
	if err := Db.Model(&article).Association("Tags").Append(t).Error; err != nil {
		return err
	}
	return nil
}
