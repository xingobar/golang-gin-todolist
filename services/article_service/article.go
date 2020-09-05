package article_service

import (
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/util"
	"golang-gin-todolist/repository"
)

type ArticleService struct {
	article *models.Article
	tag *models.Tag
	tagRepository interfaces.ITagRepository
	articleRepository interfaces.IArticleRepository
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		article: &models.Article{},
		tag: &models.Tag{},
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
	models.Db.Scopes(util.Paginate(page)).Find(&sliceArticle)
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
func (s *ArticleService) DeleteById(id string) error {
	article, err := s.articleRepository.GetById(id)
	if err != nil {
		return err
	}

	// 刪除文章
	if err := models.Db.Model(&article).Delete(&article).
		Association("Tags").
		Delete(&article.Tags).Error; err != nil {
			return err
	}
	return nil
}

func (s *ArticleService) GetByUserId(userId string) ([]models.Article, error){
	var article []models.Article
	if err := models.Db.Preload("Tags").Where("user_id = ?", userId).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}