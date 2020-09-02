package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/article_service"
	"golang-gin-todolist/services/tag_service"
	"net/http"
)

type articleController struct {
	service *article_service.ArticleService
	tagService *tag_service.TagService
}

func NewArticleController() *articleController{
	return &articleController{
		service:article_service.NewArticleService(),
		tagService: tag_service.NewTagService(),
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

	t := c.tagService.GetByIds(tags)
	if len(tags) != len(t) {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	c.service.Create(article, t)
}