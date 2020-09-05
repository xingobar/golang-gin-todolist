package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/article_service"
	"golang-gin-todolist/services/tag_service"
	validation2 "golang-gin-todolist/validation"
	article2 "golang-gin-todolist/validation/article"
	"net/http"
	"strconv"
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
	var validation article2.CreateArticleValidation
	if err := context.ShouldBind(&validation); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation2.GetError(err.(validator.ValidationErrors), article2.Message),
		})
		return
	}

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

func (c *articleController) GetPaginate(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	articles, err := c.service.GetPaginate(p)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": articles,
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

// 取得會員文章
func (c *articleController) GetByUserId (ctx *gin.Context) {

	article, err := c.service.GetByUserId(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": article,
	})
}