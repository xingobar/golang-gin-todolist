package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/resources"
	"golang-gin-todolist/services/tag_service"
	"golang-gin-todolist/validation/tag"
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

	var createTagValidation tag.CreateTagValidation
	if err := context.ShouldBind(&createTagValidation); err != nil {
		resources.NoValidResponse(context, err, tag.Message)
		return
	}

	title := context.PostForm("title")

	// 檢查標籤是否存在
	exists := t.service.ExistByName(title)
	if exists {
		resources.ErrorResponse(context, http.StatusBadRequest, e.ERROR_EXIST_TAG)
		return
	}
	ok := t.service.CreateTag(title)
	if !ok {
		resources.ErrorResponse(context, http.StatusBadRequest, e.ERROR)
		return
	}
	resources.SuccessResponse(context, e.GetMsg(e.SUCCESS))
}

// 取得所有標籤
func (t *tagController) GetAll(context *gin.Context){
	tags := t.service.GetTags()
	if tags == nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.ERROR)
		return
	}
	resources.SuccessResponse(context, tags)
}

// 取得單一標籤
func (t *tagController) GetById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.ERROR)
		return
	}

	tag := t.service.GetById(int(id))

	if tag == nil {
		resources.ErrorResponse(context, http.StatusNotFound, e.NOT_EXISTS_TAG)
		return
	}

	resources.SuccessResponse(context, tag)
}

// 刪除標籤
func (t *tagController) DeleteById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	if ok := t.service.DeleteById(id); !ok {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}
	resources.SuccessResponse(context, e.GetMsg(e.SUCCESS))
}

// 更新標籤名稱
func (t *tagController) UpdateById(context *gin.Context) {

	var v tag.CreateTagValidation
	if err := context.ShouldBind(&v); err != nil {
		resources.NoValidResponse(context, err, tag.Message)
		return
	}

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	if ok := t.service.UpdateById(id, context.PostForm("title")); !ok {
		resources.ErrorResponse(context, http.StatusBadRequest, e.INVALID_REQUEST)
		return
	}

	resources.SuccessResponse(context, e.GetMsg(e.SUCCESS))
}