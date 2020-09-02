package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/models"
	"golang-gin-todolist/services/article_service"
)

type articleController struct {
	service *article_service.ArticleService
}

func NewArticleController() *articleController{
	return &articleController{
		service:article_service.NewArticleService(),
	}
}

func (c *articleController) Create(context *gin.Context) {
	var article models.Article

	article = models.Article{
		Title: context.PostForm("title"),
		Content: context.PostForm("content"),
		UserId: 1,
	}

	tags := context.PostFormArray("tags[]")
	c.service.Create(article, tags)
}