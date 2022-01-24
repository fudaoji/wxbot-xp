package route

import (
	"wxbot-xp/controller"

	"github.com/gin-gonic/gin"
)

// 初始化消息相关路由
func initMessageRoute(app *gin.Engine) {
	group := app.Group("/message")

	// 向指定好友发送文本消息
	group.POST("/user", controller.SendTextToFriendHandle)
}
