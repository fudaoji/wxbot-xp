package controller

import (
	"wxbot-xp/bot"
	"wxbot-xp/core"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 发送消息请求体
type sendMsgRes struct {
	// username
	To string `form:"to" json:"to"`
	// 正文
	Content string `form:"content" json:"content"`
}

func SendTextToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	err := bot.SockConn.WriteMessage(websocket.TextMessage, bot.SendTxt(res.To, res.Content))
	if err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}
