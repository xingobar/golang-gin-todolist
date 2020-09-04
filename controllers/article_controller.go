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

// 新增文章
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

	if ok := c.service.Create(article, t); !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}

func (c *articleController) GetById(context *gin.Context) {
	article, err := c.service.GetById(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": article,
	})
}

// 取得所有文章
func (c *articleController) GetAll(context *gin.Context) {
	articles, err := c.service.GetAll()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": articles,
	})
}

func (c *articleController) DeleteById(context *gin.Context) {
	if err := c.service.DeleteById(context.Param("id")); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}