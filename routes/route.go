package routes

import (
	"github.com/betacraft/yaag/irisyaag"
	"github.com/kataras/iris/v12"
	"go-web-learn/controllers"
	"go-web-learn/middleware"
)

/**
服务路由注册配置
*/
func Register(api *iris.Application) {
	// api 父路由注册管理
	main := api.Party("/", middleware.CrsAuth()).AllowMethods(iris.MethodOptions)
	{

		api.Get("/", func(ctx iris.Context) { // 首页模块
			_ = ctx.View("index.html")
		})
		//注册 文档地址
		api.Get("apiDoc", func(ctx iris.Context) {
			_ = ctx.View("apiDoc/index.html")
		})

		// 注册 子二级路由
		v1 := main.Party("/pcmd/v1")
		{
			v1.PartyFunc("/", func(auth iris.Party) {
				v1.Use(irisyaag.New())
				// 设置 角色的HTTP权限控制
				//casbinMiddleware := middleware.New(model.User)

				// 设置 jwt 用户身份校验，并校验用户权限
				//auth.Use(middleware.JwtHandLer().Serve)
				auth.PartyFunc("/user", func(user iris.Party) {
					user.Post("/register", controllers.Register).Name = "用户注册"
					user.Post("/login", controllers.Login).Name = "用户登陆"
				})

			})
		}
	}

}
