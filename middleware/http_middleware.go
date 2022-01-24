package middleware

import (
	"wxbot-xp/core"

	"github.com/gin-gonic/gin"
)

// CheckAppKeyIsLoggedInMiddleware 检查AppKey是否已登录微信
func CheckAppKeyIsLoggedInMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKey := ctx.Request.Header.Get("AppKey")

		flag := true
		// 判断AppKey是否存在，商业化后根据appkey收费
		if !checkAppKey(appKey) {
			core.FailWithMessage("AppKey非法", ctx)
			flag = false
		}
		if flag {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}

// CheckAppKeyExistMiddleware 检查是否有appKey
func CheckAppKeyExistMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKey := ctx.Request.Header.Get("AppKey")
		if len(appKey) < 1 { // 先判断AppKey是不是传了
			core.FailWithMessage("AppKey为必传参数", ctx)
			ctx.Abort()
		} else if !checkAppKey(appKey) { // 判断AppKey是否存在，商业化后根据appkey收费
			core.FailWithMessage("AppKey非法", ctx)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}

//验证AppKey是否存在
func checkAppKey(appKey string) bool {
	return core.GetVal("appkey", "") == appKey
}
