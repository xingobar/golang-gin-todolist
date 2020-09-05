package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/tag_service"
	"net/http"
	"strconv"
)

type tagController struct {
	service *tag_service.TagService
}

// new tag controller
func NewTagController() *tagController {
	return &tagController{
		service: tag_service.NewTagService(),
	}
}

// 新增標籤
func (t *tagController) Create(context *gin.Context) {
	title := context.PostForm("title")

	// 檢查標籤是否存在
	exists := t.service.ExistByName(title)
	if exists {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR_EXIST_TAG,
			"msg": e.GetMsg(e.ERROR_EXIST_TAG),
		})
		return
	}
	ok := t.service.CreateTag(title)
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg": e.GetMsg(e.ERROR),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}

// 取得所有標籤
func (t *tagController) GetAll(context *gin.Context){
	tags := t.service.GetTags()
	if tags == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg": e.GetMsg(e.ERROR),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"data": tags,
	})
}

// 取得單一標籤
func (t *tagController) GetById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg": e.GetMsg(e.ERROR),
		})
		return
	}

	tag := t.service.GetById(int(id))

	if tag == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"code": e.NOT_EXISTS_TAG,
			"msg": e.GetMsg(e.NOT_EXISTS_TAG),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": tag,
	})
}

// 刪除標籤
func (t *tagController) DeleteById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	if ok := t.service.DeleteById(id); !ok {
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

// 更新標籤名稱
func (t *tagController) UpdateById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	if ok := t.service.UpdateById(id, context.PostForm("title")); !ok {
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