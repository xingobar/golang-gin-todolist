package models

import "time"

type Article struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Title string
	Content string
	UserId int
	Tags []Tag `gorm:"many2many:article_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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
func (a *Article) GetById(id string) (*Article, error) {
	var article Article
	if err := Db.Preload("Tags").Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// 取得所有文章
func (a *Article) GetAll() ([]Article, error) {
	var articles []Article
	if err := Db.Preload("Tags").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

