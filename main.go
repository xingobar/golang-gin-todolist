package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	jwt2 "golang-gin-todolist/jwt"
	"golang-gin-todolist/pkg/cache"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/router"
	"log"
	"net/http"
	"os"
)

// migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up

func main() {
	r := gin.Default()

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
	articles := r.Group("/article")
	//articles.Use(middleware.VerifyToken)
	router.ArticleRouter(articles)

	// 會員資訊
	users := r.Group("/users")
	router.UserRouter(users)

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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.TOKEN_ERROR,
			"msg": e.GetMsg(e.TOKEN_ERROR),
		})
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	// 判斷 token 是否還有效
	if ok && t.Valid {
		refreshUid, ok := claims["refresh_uid"]

		// 假設沒有 refresh uid 不給他 refresh token
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.REFRESH_UNUSED,
				"msg": e.GetMsg(e.REFRESH_UNUSED),
			})
			return
		}

		// 取得快取中的 refresh 資料
		userid, err := redis.Int(cache.Redis.Do("GET", refreshUid))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.REFRESH_EXPIRED,
				"msg": e.GetMsg(e.REFRESH_EXPIRED),
			})
			return
		}

		deleted, deleteErr := cache.Redis.Do("DEL", refreshUid)
		// 判斷是否刪除成功
		if deleteErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}

		// 判斷有刪除
		if deleted == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": e.UNAUTHORIZED,
				"msg": e.GetMsg(e.UNAUTHORIZED),
			})
			return
		}

		// 重新創建一個 token
		td ,err := jwt2.CreateJwtToken(userid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}

		// 將 token 刪近快取裡
		err = cache.SetTokenCache(userid, td)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": e.INVALID_REQUEST,
				"msg": e.GetMsg(e.INVALID_REQUEST),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg": td,
		})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.REFRESH_EXPIRED,
			"msg": e.GetMsg(e.REFRESH_EXPIRED),
		})
		return
	}
}