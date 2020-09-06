package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/comment_service"
	"golang-gin-todolist/validation"
	"golang-gin-todolist/validation/comments"
	"net/http"
	"strconv"
)

type commentController struct {
	service *comment_service.CommentService
}

func NewCommentController() *commentController {
	return &commentController{
		service: comment_service.NewCommentService(),
	}
}

// 新增留言
func (c *commentController) Create(ctx *gin.Context){

	var v comments.CreateCommentValidation

	if err := ctx.ShouldBind(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation.GetError(err.(validator.ValidationErrors), comments.Message),
		})
		return
	}

	parentId := ctx.DefaultPostForm("parent_id", "0")
	pid, err := strconv.Atoi(parentId)
	if err != nil {
		pid = 0
	}

	// 假如是子留言要判斷父留言是否存在
	if pid != 0 {
		if ok := c.service.CheckParentExists(v.ParentId); !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.PARENT_COMMENT_NOT_EXISTS,
				"msg": e.GetMsg(e.PARENT_COMMENT_NOT_EXISTS),
			})
			return
		}
	}

	// TODO: 留言的UserId 要改
	comment := models.Comment{
		Content: v.Content,
		ParentId: uint(pid),
		ArticleId: v.ArticleId,
		UserId: 1,
	}

	if ok := c.service.Create(comment); !ok {
		// 留言失敗
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	// 新增成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}