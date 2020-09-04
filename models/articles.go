package models

import "time"

type Article struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id" `
	Title string `json:"title"`
	Content string `json:"content"`
	UserId int	`json:"user_id"`
	Tags []Tag `gorm:"many2many:article_tags;" json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
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

