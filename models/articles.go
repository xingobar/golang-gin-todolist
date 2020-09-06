package models

import "time"

type Article struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id" `
	Title string `json:"title"`
	Content string `json:"content"`
	UserId int	`json:"user_id"`
	Tags []Tag `gorm:"many2many:article_tags;" json:"tags"`
	Comments []Comment `gorm:"foreignKey:ArticleId" json:"comments"`
	User User `gorm:"foreignKey:ID" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (Article) TableName() string{
	return "articles"
}

