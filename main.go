package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"tp_iris/log"
	"tp_iris/config"
	"tp_iris/router"
)

func main() {
	// 基于router实例化 iris app 对象
	app := router.SetupRouter()

	// 开启 iris accesslog 中间件功能
	ac := log.MakeAccessLog(); defer ac.Close()
	app.UseGlobal(ac.Handler)

	// 自定义配置项
	fmt.Println(config.C.MySQL.User, config.C.MySQL.Passwd, &config.C.Log.Ac)

	// RUN
	addr := config.C.Iris.Other["addr"].(string)
	_ = app.Listen(addr, iris.WithConfiguration(config.C.Iris))
}

