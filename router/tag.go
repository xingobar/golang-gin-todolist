package router

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/controllers"
)

func TagRouter(group *gin.RouterGroup) {
	tagController := controllers.NewTagController()

	// 新增標籤
	group.POST("/create", tagController.Create)

	// 取得所有標籤
	group.GET("/", tagController.GetAll)

	// 取得單一標籤
	group.GET("/edit/:id", tagController.GetById)

	// 刪除單一標籤
	group.DELETE("/delete/:id", tagController.DeleteById)

	// 更新標籤名稱
	group.PUT("/update/:id", tagController.UpdateById)
}