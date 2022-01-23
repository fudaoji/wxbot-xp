package route

import (
	"wxbot-xp/controller"

	"github.com/gin-gonic/gin"
)

// initUserRoute 初始化登录路由信息
func initUserRoute(app *gin.Engine) {
	group := app.Group("/user")
	// hello
	group.GET("/hello", controller.Hello)
}
