package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/jwt"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/resources"
	"golang-gin-todolist/services/article_service"
	"golang-gin-todolist/services/tag_service"
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
		resources.NoValidResponse(context, err,  article2.Message)
		return
	}

	// 取得 token 的 userid 資訊
	accessDetail, err := jwt.ExtractTokenMetadata(context.Request)
	if err != nil {
		resources.ErrorResponse(context, http.StatusUnauthorized, e.UNAUTHORIZED)
		return
	}

	var article models.Article
	article = models.Article{
		Title: context.PostForm("title"),
		Content: context.PostForm("content"),
		UserId: accessDetail.UserId,
	}

	tags := context.PostFormArray("tags[]")

	t := c.tagService.GetByIds(tags)
	if len(tags) != len(t) {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	if ok := c.service.Create(article, t); !ok {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}
	resources.SuccessResponse(context, e.GetMsg(e.SUCCESS))
}

func (c *articleController) GetById(context *gin.Context) {
	article, err := c.service.GetById(context.Param("id"))

	if err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	resources.SuccessResponse(context, article)
}

func (c *articleController) GetPaginate(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	p, err := strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	articles, err := c.service.GetPaginate(p)

	resources.SuccessResponse(ctx, articles)
}

// 取得所有文章
func (c *articleController) GetAll(context *gin.Context) {
	articles, err := c.service.GetAll()

	if err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	resources.SuccessResponse(context, articles)
}

func (c *articleController) DeleteById(context *gin.Context) {
	accessDetail,err := jwt.ExtractTokenMetadata(context.Request)
	if err != nil {
		resources.ErrorResponse(context, http.StatusUnauthorized, e.TOKEN_ERROR)
		return
	}
	if err := c.service.DeleteById(accessDetail.UserId ,context.Param("id")); err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	resources.SuccessResponse(context, e.GetMsg(e.SUCCESS))
}
