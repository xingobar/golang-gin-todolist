package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/tag_service"
	"net/http"
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
	err := t.service.CreateTag(title)
	if err != nil {
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