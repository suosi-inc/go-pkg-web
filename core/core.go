package core

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ClientIpByHeader(ctx *gin.Context, header string) string {
	realIp := strings.TrimSpace(ctx.GetHeader(header))
	if realIp == "" {
		if ctx.ClientIP() == "::1" {
			return "127.0.0.1"
		} else {
			return ctx.ClientIP()
		}
	}

	return realIp
}

func ClientIp(ctx *gin.Context) string {
	return ClientIpByHeader(ctx, "X-Real-IP")
}

func ClientUserAgent(ctx *gin.Context) string {
	userAgent := strings.TrimSpace(ctx.GetHeader("user-agent"))
	if userAgent != "" {
		return userAgent
	}

	return ""
}
