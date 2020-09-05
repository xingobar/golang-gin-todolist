package article_service

import (
	"fmt"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/setting"
	"golang-gin-todolist/pkg/util"
	"math"
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

// 根據編號取得文章
func (s *ArticleService) GetById(id string) (*models.Article, error) {
	article, err := s.article.GetById(id)
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
	perPage := setting.TWENTY_PAGE
	totalPage := int(math.Ceil(float64(len(articles) / perPage)))
	start := (page - 1) * perPage
	last := perPage * page

	if page > totalPage {
		start = (totalPage - 2) * perPage
		last =  (totalPage - 1) * perPage
	}

	paginator := new(util.Paginator)
	paginator.Total = len(articles)
	paginator.Data = articles[start: last]
	paginator.TotalPage = int(math.Ceil(float64(len(articles) / perPage)))
	paginator.Page = page

	return paginator, nil
}

// 取得所有文章
func (s *ArticleService) GetAll() ([]models.Article, error) {
	articles, err  := s.article.GetAll()
	fmt.Println("======== length =======")
	fmt.Println(len(articles))
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// 根據編號刪除文章
func (s *ArticleService) DeleteById(id string) error {
	article, err := s.article.GetById(id)
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