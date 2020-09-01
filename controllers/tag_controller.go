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