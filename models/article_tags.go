package models

type ArticleTags struct {
	ID int `gorm:"primaryKey;autoIncrement;column:id"`
	ArticleId int `gorm:"column:article_id"`
	TagId	int `gorm:"column:tag_id"`
}

func (ArticleTags) TableName() string {
	return "article_tags"
}