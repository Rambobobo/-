package main

import (
	"github.com/kataras/iris/v12"
	"myweb/controller"
)

func main() {
	//app := iris.New()
	//
	//userController := &controller.UserController{}
	//app.Post("/login", userController.Login)
	//app.Run(iris.Addr(":9001"))
	//mvc路由组优化后

	app := newApp()
	//路由设置
	mvcHadnle(app)
	app.Run(
		iris.Addr(":9001"),     //监听9001端口
		iris.WithOptimizations, //对json数据序列化更快的配置
	)
}

// 构建app
func newApp() *iris.Application {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	return app
}

// mvc
func mvcHadnle(app *iris.Application) {
	//创建控制器对象
	userController := &controller.UserController{}
	//用户路由组
	userRouter := app.Party("/user")
	{
		userRouter.Post("/login", userController.Login)
		userRouter.Post("/register", userController.Register)
	}
}
