package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/validation"
	"net/http"
)

// 錯誤訊息
func ErrorResponse(ctx *gin.Context, httpStatus int, code int) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})
}

// 成功
func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": data,
	})
}

// 驗證錯誤
func NoValidResponse(ctx *gin.Context, err error, msg map[string]string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": e.INVALID_REQUEST,
		"msg": validation.GetError(err.(validator.ValidationErrors), msg),
	})
}