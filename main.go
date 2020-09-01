package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
	"golang-gin-todolist/services/tag_service"
	"net/http"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()

	tagController := controllers.NewTagController()
	r.POST("/create", tagController.Create)

	// 標籤
	tags := r.Group("/tags")
	{
		// 新增標籤
		tags.POST("/create", func(context *gin.Context) {
			tag := tag_service.Tag{Title:context.PostForm("title")}
			err := tag.AddTag()
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status": "add successful",
			})
		})
	}

	r.Run()
}
