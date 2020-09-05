package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

func ArticleRouter(router *gin.RouterGroup) {
	articleController := controllers.NewArticleController()
	// 新增文章
	router.POST("/create", articleController.Create)

	// 取得文章
	router.GET("/edit/:id", articleController.GetById)

	// 取得所有文章
	router.GET("/", articleController.GetAll)

	// 刪除文章
	router.DELETE("/:id", articleController.DeleteById)
}