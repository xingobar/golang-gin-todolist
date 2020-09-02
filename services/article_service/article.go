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
	if err := s.article.Create(article, tags); err != nil {
		return false
	}
	return true
}
