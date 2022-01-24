package main

import (
	"wxbot-xp/bot"
	"wxbot-xp/core"
	"wxbot-xp/middleware"

	"github.com/gin-gonic/gin"

	"wxbot-xp/logger"

	"wxbot-xp/route"
)

// 程序启动入口
func main() {
	// 初始化日志
	logger.InitLogger()
	gin.SetMode(gin.ReleaseMode)
	// 初始化Gin
	app := gin.Default()
	// 定义全局异常处理
	app.NoRoute(core.NotFoundErrorHandler())
	// AppKey预检
	app.Use(middleware.CheckAppKeyExistMiddleware(), middleware.CheckAppKeyIsLoggedInMiddleware())

	// 初始化路由
	route.InitRoute(app)
	//读取配置
	core.InitConfig()
	//启动websocket
	go bot.InitWSConnHandle()

	// 监听端口
	_ = app.Run(":8889")
}
