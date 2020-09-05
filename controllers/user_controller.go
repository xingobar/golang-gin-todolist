package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang-gin-todolist/jwt"
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

	if err := ctx.ShouldBind(&v); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation.GetError(err.(validator.ValidationErrors), user.Message),
		})
		return
	}

	// TODO: 重構
	var user  = models.User{
		Username: v.Username,
		Email: v.Email,
		Password: v.Password,
	}

	// 檢查帳號是否存在
	ok, err := c.service.CheckExistByEmail(v.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}

	// 帳號存在
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.EXISTS_EMAIL,
			"msg": e.GetMsg(e.EXISTS_EMAIL),
		})
		return
	}

	// 註冊
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

// 登入
func (c *userController) Login(ctx *gin.Context) {
	var login user.LoginValidation

	if err := ctx.ShouldBind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": validation.GetError(err.(validator.ValidationErrors), user.Message),
		})
		return
	}

	user, err := c.service.GetUserByEmailAndPassword(login.Email, login.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.LOGIN_ERROR,
			"msg": e.GetMsg(e.LOGIN_ERROR),
		})
		return
	}

	id := int(user.ID)
	token, err := jwt.CreateJwtToken(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.TOKEN_ERROR,
			"msg": e.GetMsg(e.TOKEN_ERROR),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": token,
	})
}