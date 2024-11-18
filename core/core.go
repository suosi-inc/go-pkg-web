package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ClientIp(ctx *gin.Context) string {
	// Nginx 反向代理
	realIp := strings.TrimSpace(ctx.GetHeader("X-Real-IP"))
	if realIp == "" {
		if ctx.ClientIP() == "::1" {
			return "127.0.0.1"
		} else {
			return ctx.ClientIP()
		}
	}

	return realIp
}

func ClientUserAgent(ctx *gin.Context) string {
	userAgent := strings.TrimSpace(ctx.GetHeader("user-agent"))
	if userAgent != "" {
		return userAgent
	}

	return ""
}
