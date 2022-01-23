package controller

import (
	"wxbot-xp/core"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	core.OkWithMessage("success", ctx)
}
