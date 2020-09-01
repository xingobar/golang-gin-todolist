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
	}

	r.Run()
}
