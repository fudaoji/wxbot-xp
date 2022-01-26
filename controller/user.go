package controller

import (
	"wxbot-xp/bot"
	"wxbot-xp/core"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetCurrentUserInfoHandle 获取当前登录用户
func GetCurrentUserInfoHandle(ctx *gin.Context) {
	err := bot.GetInstance().WriteMessage(websocket.TextMessage, bot.GetPersonalInfo())
	if err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}

	// 返回给前端
	core.Ok(ctx)
}

// GetFriendsListHandle 获取好友列表
func GetFriendsListHandle(ctx *gin.Context) {
	err := bot.GetInstance().WriteMessage(websocket.TextMessage, bot.GetContactList())
	if err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}

	// 返回给前端
	core.Ok(ctx)
}

// GetGroupsListHandle 获取群聊列表
func GetGroupsListHandle(ctx *gin.Context) {

}

func Hello(ctx *gin.Context) {
	core.OkWithMessage("success", ctx)
}
