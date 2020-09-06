package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

func CommentRouter(group *gin.RouterGroup) {
	commentController := controllers.NewCommentController()

	// 新增留言
	group.POST("/", commentController.Create)
}
