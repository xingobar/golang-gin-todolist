package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

func CommentRouter(group *gin.RouterGroup) {
	commentController := controllers.NewCommentController()

	// 新增留言
	group.POST("/", commentController.Create)

	// 取得子留言
	group.GET("/:id", commentController.GetChildCommentById)
}
