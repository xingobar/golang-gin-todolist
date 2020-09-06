package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/jwt"
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

	// 父留言編號
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


	accessDetail, err := jwt.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": e.UNAUTHORIZED,
			"msg": e.GetMsg(e.UNAUTHORIZED),
		})
		return
	}

	// TODO: 留言的UserId 要改
	//var comment models.Comment
	//b := new(bytes.Buffer)
	//json.NewEncoder(b).Encode(v)
	//err = json.Unmarshal(b.Bytes(), &comment)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"code": e.INVALID_REQUEST,
	//		"msg": e.GetMsg(e.INVALID_REQUEST),
	//	})
	//	return
	//}
	//comment.UserId = 1
	comment := models.Comment{
		Content: v.Content,
		ParentId: pid,
		ArticleId: v.ArticleId,
		UserId: accessDetail.UserId,
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

// 取得子留言
func (c *commentController) GetChildCommentById(ctx *gin.Context) {
	id := ctx.Param("id")
	parentId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	if ok := c.service.CheckParentExists(uint(parentId)); !ok {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": e.PARENT_COMMENT_NOT_EXISTS,
			"msg": e.GetMsg(e.PARENT_COMMENT_NOT_EXISTS),
		})
		return
	}

	data, err := c.service.GetChildComment(uint(parentId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": data,
	})
}

func (c *commentController) DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	value, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	accessDetail, err := jwt.ExtractTokenMetadata(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.TOKEN_ERROR,
			"msg": e.GetMsg(e.TOKEN_ERROR),
		})
		return
	}

	comment, err := c.service.GetById(value)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": e.NOT_FOUND,
			"msg": e.GetMsg(e.NOT_FOUND),
		})
		return
	}

	// 不是該會員的文章
	if accessDetail.UserId != comment.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": e.UNAUTHORIZED,
			"msg": e.GetMsg(e.UNAUTHORIZED),
		})
		return
	}

	if comment.ParentId == 0 {
		_, err := c.service.DeleteParentById(*comment)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}
	} else {
		// 單純子留言
		ok, err := c.service.DeleteChildById(int(comment.ID))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}

		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}