package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
	"golang-gin-todolist/middleware"
)

func CommentRouter(group *gin.RouterGroup) {
	commentController := controllers.NewCommentController()

	// 新增留言
	group.POST("/", commentController.Create).Use(middleware.VerifyToken)

	// 取得子留言
	group.GET("/:id", commentController.GetChildCommentById)
}
