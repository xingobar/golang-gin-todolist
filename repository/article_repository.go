package repository

import (
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
)

type articleRepository struct {

}

func NewArticleRepository() interfaces.IArticleRepository {
	return &articleRepository{}
}

// 新增文章
func (repository *articleRepository) Create(article models.Article, tags []models.Tag) (error) {
	tx := models.Db.Begin()
	tx.Create(&article)
	if err := tx.Model(&article).Association("Tags").Append(tags).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 取得單一文章
func (repository *articleRepository) GetById(id string) (*models.Article, error) {
	var article models.Article
	if err := models.Db.Preload("Tags").Preload("User").Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (repository *articleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	if err := models.Db.Preload("Tags").Preload("User").Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
