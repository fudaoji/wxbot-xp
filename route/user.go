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

	// 获取好友列表
	group.GET("/friends", controller.GetFriendsListHandle)
	// 获取当前用户信息
	group.GET("/info", controller.GetCurrentUserInfoHandle)
}
