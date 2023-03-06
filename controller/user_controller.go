package controller

import (
	"github.com/kataras/iris/v12"
	"myweb/model"
	"myweb/service"
)

type UserController struct{}

func (uc *UserController) Login(ctx iris.Context) {
	/*json请求*/
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//请求格式检测
	if err := ctx.ReadJSON(&loginRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "登录失败",
			"message": "请求格式错误",
		})
		return
	}

	userService := &service.UserService{}
	loggedUser, err := userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "登录失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "1",
		"success": "登录成功",
		"message": "用户登录成功",
		"user":    loggedUser,
	})
}
func (uc *UserController) Register(ctx iris.Context) {
	/*json请求*/
	var RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//请求格式检测
	if err := ctx.ReadJSON(&RegisterRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "登录失败",
			"message": "请求格式错误",
		})
		return
	}
	userService := &service.UserService{}
	user := &model.User{
		UserName: RegisterRequest.Username,
		Password: RegisterRequest.Password,
	}
	err := userService.Register(user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "0",
			"success": "注册失败",
			"message": err.Error(),
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"status":  "1",
		"success": "注册成功",
		"message": "注册成功",
		"user":    user,
	})
}
