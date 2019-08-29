package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func Cors() context.Handler {
	return func(ctx iris.Context) {
		// 处理js-ajax跨域问题
		ctx.Header("Access-Control-Allow-Origin", "*") //允许访问所有域
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS, POST")
		ctx.Header("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,X-Auth-Token")
		ctx.Header("Access-Control-Allow-Headers", "Access-Token")
		//放行所有OPTIONS方法
		method := ctx.Method
		if method() == "OPTIONS" {
			ctx.JSON(iris.Map{
				"code": 1,
				"msg":  "跨域",
			})
		}
		ctx.Next()
	}
}
