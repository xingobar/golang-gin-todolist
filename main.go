package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
	"golang-gin-todolist/router"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()

	//r.GET("/run", func(context *gin.Context) {
	//	fmt.Println("test")
	//	models.Db.AutoMigrate(&models.Article{}, &models.Tag{})
	//	//models.Db.CreateTable(&models.Tag{})
	//	fmt.Printf("test")
	//})

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
	//articles.Use(middleware.VerifyToken)
	router.ArticleRouter(articles)

	// 會員資訊
	users := r.Group("/users")
	router.UserRouter(users)

	r.Run()
}
