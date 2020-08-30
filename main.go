package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()

	// 標籤
	r.Group("/tags")
	{
		// 新增標籤
		r.POST("/create", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"status": "add successful",
			})
		})
	}
}
