package article_service

import (
	"fmt"
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

func (s *ArticleService) Create(article models.Article, tags []models.Tag) bool {

	if err := s.article.Create(article, tags); err != nil {
		fmt.Println(err.Error())
		return false
	}

	//fmt.Println("id: ", article.ID)
	////tag, _ := s.tag.GetById(1)
	//err := s.article.AddTags(&article, models.Tag{})
	//fmt.Println("err =====", err)

	//if err := s.article.AddTags(&article, tags); err != nil {
	//	fmt.Println(err.Error())
	//	return false
	//}
	//fmt.Println("===============")
	//fmt.Println(article.ID)
	//fmt.Println(s.article.AddTags(a, tags))

	return true
}
