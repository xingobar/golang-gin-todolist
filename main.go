package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()



	// 標籤
	tags := r.Group("/tags")
	{
		tagController := controllers.NewTagController()

		// 新增標籤
		tags.POST("/create", tagController.Create)

		// 取得所有標籤
		tags.GET("/", tagController.GetAll)

		// 取得單一標籤
		tags.GET("/edit/:id", tagController.GetById)

		// 刪除單一標籤
		tags.DELETE("/delete/:id", tagController.DeleteById)

		// 更新標籤名稱
		tags.PUT("/update/:id", tagController.UpdateById)
	}

	// 文章
	articles := r.Group("/article")
	{
		articleController := controllers.NewArticleController()
		// 新增文章
		articles.POST("/create", articleController.Create)

		// 取得文章
		articles.GET("/edit/:id", articleController.GetById)

		// 取得所有文章
		articles.GET("/", articleController.GetAll)
	}

	r.Run()
}
