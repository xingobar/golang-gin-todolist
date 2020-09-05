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
	router.POST("/register", userController.Register)

	// 登入
	router.POST("/login", userController.Login)
}