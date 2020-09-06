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

// 取得子留言
func (c *CommentService) GetChildComment(id uint) ([]models.Comment, error) {
	var comments []models.Comment

	if err := models.Db.Where("parent_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// 取得留言
func (c *CommentService) GetById(id int) (*models.Comment, error){
	var comment models.Comment
	if err := models.Db.Where("id = ? ", id).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// 刪除子留言
func (c *CommentService) DeleteChildById(id int) (bool, error) {
	if err := models.Db.Where("id = ?", id).Delete(&models.Comment{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// 刪除文章
func (c *CommentService) DeleteParentById(comment models.Comment) (bool, error) {
	if err := models.Db.Where("id = ? or parent_id = ?",
								comment.ID,
								comment.ID).
				Delete(&models.Comment{}).
				Delete(&models.Comment{}).
				Error; err != nil {
		return false, err
	}
	return true, nil
}