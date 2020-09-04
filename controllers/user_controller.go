package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/validation"
	"golang-gin-todolist/validation/user"
	"net/http"
)

type userController struct {

}

func NewUserController() *userController{
	return &userController{}
}

// 註冊
func (c *userController) Register(ctx *gin.Context) {

	var v user.RegisterValidation

	//var user models.User
	if err := ctx.ShouldBind(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation.GetError(err.(validator.ValidationErrors), user.Message),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}