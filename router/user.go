package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

func UserRouter(router *gin.RouterGroup) {
	articleController := controllers.NewArticleController()
	userController := controllers.NewUserController()

	// 會員文章
	router.GET("/article/:id", articleController.GetByUserId)

	// 註冊
	router.POST("/", userController.Register)
}