package main

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
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
	{
		articleController := controllers.NewArticleController()
		// 新增文章
		articles.POST("/create", articleController.Create)

		// 取得文章
		articles.GET("/edit/:id", articleController.GetById)

		// 取得所有文章
		articles.GET("/", articleController.GetAll)

		articles.DELETE("/:id", articleController.DeleteById)
	}

	// 會員資訊
	users := r.Group("/users")
	{
		articleController := controllers.NewArticleController()
		userController := controllers.NewUserController()

		// 會員文章
		users.GET("/article/:id", articleController.GetByUserId)

		// 註冊
		users.POST("/", userController.Register)
	}

	r.Run()
}
