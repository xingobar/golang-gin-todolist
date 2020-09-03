package article_service

import (
	"golang-gin-todolist/models"
)

type ArticleService struct {
	article *models.Article
	tag *models.Tag
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		article: &models.Article{},
		tag: &models.Tag{},
	}
}

// 新增文章
func (s *ArticleService) Create(article models.Article, tags []models.Tag) bool {
	tx := models.Db.Begin()
	tx.Create(&article)
	if err := tx.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func (s *ArticleService) GetById(id string) *models.Article {
	article := s.article.GetById(id)
	if article == nil {
		return nil
	}

	var tags []models.Tag
	if err := models.Db.Model(&article).Association("Tags").Find(&tags).Error; err != nil {
		return nil;
	}
	article.Tags = tags
	return article
}
