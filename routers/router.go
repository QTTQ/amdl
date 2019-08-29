// Package routers provides ...
package routers

import (
	"amdl/controllers"
	"amdl/middleware"
	"fmt"

	// "github.com/betacraft/yaag/yaag"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	// "github.com/kataras/iris/core/router"
)

func AuthTokenHandler(ctx iris.Context) {
	fmt.Println("aaaa:aaaaa")
	user := ctx.Values().Get("token").(*jwt.Token)
	fmt.Println(user, ":aaaaa")
	if user != nil {
		ctx.Next()
	}
	ctx.Next()

}

// InitRouters aaa
func InitRouters(router *iris.Application) *iris.Application {
	//app.RegisterView(iris.HTML("./web", ".html")) //加载模版文件
	//  app.StaticWeb("/static", "web/resources/static") // 设置静态资源,暂时没有
	// app.RegisterView(iris.HTML("web/views", ".html").Reload(true))
	// router.Use(middleware.Cors())
	// router.Use(middleware.AllUrlPath)
	router.Use(middleware.Cors())
	router.Post("/Register", controllers.Register) // 注册
	router.Post("/LoginIn", controllers.LoginIn)   //登录
	// /api/recommendedDisplayPictures
	apiRouter := router.Party("/api")
	{
		// 推荐展示图片
		apiRouter.Get("/recommendedDisplayPictures", controllers.RecommendedDisplayPictures)

		// 单个人的社区详情
		apiRouter.Get("/service/personalCommunityInfo", controllers.PersonalCommunityInfo)
	}

	jwtRouter := router.Party("/jwt")
	jwtRouter.Use(middleware.UserAuth())
	{
		jwtRouter.Post("/aaa", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"code": "aaaa",
			})
		}) // 注册
	}

	return router
}
