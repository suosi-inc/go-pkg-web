package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/suosi-inc/go-webpkg/middleware/jwts"
	"github.com/suosi-inc/go-webpkg/rest"
)

// JwtAuth jwt 中间件
func JwtAuth(headerKey string, secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(headerKey)
		if token == "" {
			rest.JsonUnauthorizedMessage(ctx, "jwt auth")
			ctx.Abort()
		} else {
			claims, err := jwts.ParseToken(token, secret)
			if err != nil {
				rest.JsonUnauthorizedMessage(ctx, err.Error())
				ctx.Abort()
				return
			} else {
				ctx.Set("Jwt-Claims", claims)
			}
		}

		ctx.Next()
	}
}
