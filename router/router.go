package router

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/rs/cors"
	"tp_iris/controller"
)

var Router = iris.New()

func init() {
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	Router.UseRouter(recover.New())
}


func SetupRouter() *iris.Application {
	// CORS 跨域资源共享
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	})

	// of the same form of app.WrapRouter's accept input argument,
	Router.WrapRouter(crs.ServeHTTP)

	v1 := Router.Party("/v1").AllowMethods(iris.MethodOptions) // <- 对于预检很重要。
	{
		v1.Get("/info", controller.Info)
	}

	return Router
}