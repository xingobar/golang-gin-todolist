package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/ulule/deepcopier"
	"golang-gin-todolist/jwt"
	"golang-gin-todolist/models"
	"golang-gin-todolist/pkg/cache"
	"golang-gin-todolist/pkg/e"
	"golang-gin-todolist/pkg/util"
	resources2 "golang-gin-todolist/resources"
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
		Password: util.HashAndSalt([]byte(v.Password)),
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
	if ok := c.service.Register(user); !ok {
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

	user, err := c.service.GetUserByEmail(login.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.LOGIN_ERROR,
			"msg": e.GetMsg(e.LOGIN_ERROR),
		})
		return
	}

	// 檢查密碼是否一致
	if ok := util.ComparePasswords(user.Password, []byte(login.Password)); !ok {
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

	// 把token 存到快取
	saveErr := cache.SetTokenCache(id, token)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.CACHE_ERROR,
			"msg": e.GetMsg(e.CACHE_ERROR),
		})
		return
	}

	resources := &resources2.UserResource{}
	deepcopier.Copy(token).To(resources)
	deepcopier.Copy(user).To(resources)

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": resources,
	})
}

// 取得會員文章
func (c *userController) GetArticles (ctx *gin.Context) {

	article, err := c.service.GetArticles(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_REQUEST,
			"msg": e.GetMsg(e.INVALID_REQUEST),
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg": article,
	})
}