package main

import (
	"net/http"
	"time"

	"amdl/config"
	"amdl/db"
	"amdl/models"
	"amdl/routers"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type User1 struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

const refreshEvery = 10 * time.Second

func main() {
	db.ConnectAndInit(
		config.Conf,
		new(models.Article),
		new(models.Comments),
		new(models.Photos),
		new(models.RecomPictures),
		new(models.User),
		new(models.Follow),
	)
	// defer db.DB.Close()
	iris.RegisterOnInterrupt(func() {
		db.DB.Close()
	})

	app := iris.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(iris.Cache304(refreshEvery))
	app.Get("*", cache.Handler(10*time.Second), func(ctx iris.Context) {})

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*", "Content-Type", "Content-Length", "Accept", "X-Requested-With", "Origin", "Authorization"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: true,
	// })

	v1 := app.Party("/", crs).AllowMethods(iris.MethodOptions) // <- 对于预检很重要。
	{

		v1.Get("/abcd", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"aaa": "成功111",
			})
		}) //登录
	}

	// crs := cors.New(cors.Options{
	// 	// AllowedOrigins:   []string{"http://foo.com"},   //允许通过的主机名称
	// 	AllowedOrigins:   []string{"*"}, //允许通过的主机名称
	// 	AllowCredentials: true,
	// })
	app.Post("/user", func(ctx iris.Context) {
		c := &User1{}
		if err := ctx.ReadForm(c); err != nil {
			ctx.JSON(iris.Map{
				"AAA": "AAA",
			})
		} else {
			ctx.JSON(c)
		}
	})
	router := routers.InitRouters(app)
	h := app.NewHost(&http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})
	h.RegisterOnShutdown(func() {
		println("server terminated")
	})
	app.Run(iris.Raw(h.ListenAndServe))
}
