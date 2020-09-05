package interfaces

import "golang-gin-todolist/models"

type IArticleRepository interface {
	// 新增文章
	Create(article models.Article, tags []models.Tag) (error)

	// 取得單一文章
	GetById(id string) (*models.Article, error)

	// 取得所有文章
	GetAll() ([]models.Article, error)

}