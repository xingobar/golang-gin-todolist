package middleware

import (
	"github.com/gin-gonic/gin"
	"golang-gin-todolist/jwt"
	"golang-gin-todolist/pkg/e"
	"net/http"
)

// 驗證 token
func VerifyToken(ctx *gin.Context) {
	_, err := jwt.ParseToken(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": e.UNAUTHORIZED,
			"msg": e.GetMsg(e.UNAUTHORIZED),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}