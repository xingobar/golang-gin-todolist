package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	jwt2 "golang-gin-todolist/jwt"
	"golang-gin-todolist/middleware"
	"golang-gin-todolist/pkg/cache"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/resources"
	"golang-gin-todolist/router"
	"io"
	"log"
	"net/http"
	"os"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerToFile)

	// 設定log
	f, _ := os.Create("./log/info.log")
	e, _ := os.Create("./log/error.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(e, os.Stdout)


	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//r.GET("/run", func(context *gin.Context) {
	//	fmt.Println("test")
	//	models.Db.AutoMigrate(&models.Article{}, &models.Tag{})
	//	//models.Db.CreateTable(&models.Tag{})
	//	fmt.Printf("test")
	//})

	// 標籤
	tags := r.Group("/tags")
	router.TagRouter(tags)

	// 文章
	// 要授權才能對文章操作
	router.AuthArticleRouter(r.Group("/article"))

	// 不用授權就可以讀取
	router.ArticleRouter(r.Group("/article"))

	// 會員資訊
	users := r.Group("/users")
	router.UserRouter(users)

	// 留言
	router.CommentRouter(r.Group("/comments"))
	router.AuthCommentRouter(r.Group("/comments"))

	// refresh 重新取得 token
	r.POST("/refresh", Refresh)

	r.Run()
}

// token refresh
func Refresh(ctx *gin.Context){
	refreshToken := ctx.PostForm("refresh_token")

	// 解碼 claims
	t, err := jwt.Parse(refreshToken, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,  fmt.Errorf("unexpected token")
		}
		return []byte(os.Getenv("REFRESH_JWT")), nil
	})

	// 解碼失敗
	if err != nil {
		resources.ErrorResponse(ctx, http.StatusBadRequest, e.TOKEN_ERROR)
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	// 判斷 token 是否還有效
	if ok && t.Valid {
		refreshUid, ok := claims["refresh_uid"]

		// 假設沒有 refresh uid 不給他 refresh token
		if !ok {
			resources.ErrorResponse(ctx, http.StatusBadRequest, e.REFRESH_UNUSED)
			return
		}

		// 取得快取中的 refresh 資料
		userid, err := redis.Int(cache.Redis.Do("GET", refreshUid))
		if err != nil {
			resources.ErrorResponse(ctx, http.StatusBadRequest, e.REFRESH_EXPIRED)
			return
		}

		deleted, deleteErr := cache.Redis.Do("DEL", refreshUid)
		// 判斷是否刪除成功
		if deleteErr != nil {
			resources.ErrorResponse(ctx, http.StatusBadRequest, e.INVALID_REQUEST)
			return
		}

		// 判斷有刪除
		if deleted == 0 {
			resources.ErrorResponse(ctx, http.StatusUnauthorized, e.UNAUTHORIZED)
			return
		}

		// 重新創建一個 token
		td ,err := jwt2.CreateJwtToken(userid)
		if err != nil {
			resources.ErrorResponse(ctx, http.StatusBadRequest, e.INVALID_REQUEST)
			return
		}

		// 將 token 刪近快取裡
		err = cache.SetTokenCache(userid, td)
		if err != nil {
			resources.ErrorResponse(ctx, http.StatusBadRequest, e.INVALID_REQUEST)
			return
		}

		resources.SuccessResponse(ctx, td)
		return
	} else {
		resources.ErrorResponse(ctx, http.StatusBadRequest, e.REFRESH_EXPIRED)
		return
	}
}