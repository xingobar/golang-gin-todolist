package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
	"golang-gin-todolist/middleware"
)

func CommentRouter(group *gin.RouterGroup) {
	commentController := controllers.NewCommentController()
	// 取得子留言
	group.GET("/:id", commentController.GetChildCommentById)
}

// 要授權才能對操作留言
func AuthCommentRouter(group *gin.RouterGroup) {
	commentController := controllers.NewCommentController()

	group.Use(middleware.VerifyToken)

	// 新增留言
	group.POST("/", commentController.Create)

	// 刪除留言
	group.DELETE("/:id", commentController.DeleteById)
}