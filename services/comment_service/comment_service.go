package comment_service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang-gin-todolist/models"
)

type CommentService struct {

}

func NewCommentService() *CommentService {
	return &CommentService{}
}

// 檢查留言是否存在
func (c *CommentService) CheckParentExists(id uint) bool{
	var comment models.Comment
	if err := models.Db.Where("id = ? ", id).First(&comment).Error; err != nil {
		// 找不到父留言
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

// 新增留言
func (c *CommentService) Create(comment models.Comment) bool{
	if err := models.Db.Create(&comment).Error; err != nil {
		return false
	}
	return true
}