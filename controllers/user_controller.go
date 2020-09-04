package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/services/user_service"
	"golang-gin-todolist/validation"
	"golang-gin-todolist/validation/user"
	"net/http"
)

type userController struct {
	service *user_service.UserService
}

func NewUserController() *userController{
	return &userController{
		service: user_service.NewUserService(),
	}
}

// 註冊
func (c *userController) Register(ctx *gin.Context) {

	var v user.RegisterValidation

	//var user models.User
	fmt.Println(ctx.PostForm("confirm_password"))
	fmt.Println(ctx.PostForm("password"))
	if err := ctx.ShouldBind(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation.GetError(err.(validator.ValidationErrors), user.Message),
		})
		return
	}

	var user  = models.User{
		Username: v.Username,
		Email: v.Email,
		Password: v.Password,
	}


	if err := c.service.Register(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": e.GetMsg(e.SUCCESS),
	})
}