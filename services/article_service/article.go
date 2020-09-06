package article_service

import (
	"fmt"
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/util"
	"golang-gin-todolist/repository"
)

type ArticleService struct {
	tagRepository interfaces.ITagRepository
	articleRepository interfaces.IArticleRepository
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		tagRepository: repository.NewTagRepository(),
		articleRepository: repository.NewArticleRepository(),
	}
}

// 新增文章
func (s *ArticleService) Create(article models.Article, tags []models.Tag) bool {
	if err := s.articleRepository.Create(article, tags); err != nil {
		return false
	}
	return true
}

// 根據編號取得文章
func (s *ArticleService) GetById(id string) (*models.Article, error) {
	article, err := s.articleRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// 取得分頁
func (s *ArticleService) GetPaginate(page int) (*util.Paginator, error){
	var articles []models.Article
	if err := models.Db.Preload("Tags").Find(&articles).Error; err != nil {
		return nil, err
	}
	var sliceArticle []models.Article
	models.Db.Scopes(util.Paginate(page)).Preload("Tags").Preload("User").Find(&sliceArticle)
	return util.CreatePaginate(len(articles), sliceArticle, page), nil
}

// 取得所有文章
func (s *ArticleService) GetAll() ([]models.Article, error) {
	articles, err  := s.articleRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// 根據編號刪除文章
func (s *ArticleService) DeleteById(userid int, id string) error {
	article, err := s.articleRepository.GetById(id)
	if err != nil {
		return err
	}

	// 判斷是不是你的文章
	if article.UserId != userid {
		return fmt.Errorf("not your article")
	}

	// 刪除文章
	if err := models.Db.Model(&article).Delete(&article).
		Association("Tags").
		Delete(&article.Tags).Error; err != nil {
			return err
	}
	return nil
}