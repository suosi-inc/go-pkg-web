package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CorsDefaultAllowOrigin      = "*"
	CorsDefaultAllowMethods     = "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS"
	CorsDefaultAllowHeaders     = "Origin, Content-Type, X-CSRF-Token, Authorization, Jwt-Authorization, X-Requested-With, Accept, X-Cookie"
	CorsDefaultExposeHeaders    = "Content-Length, Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma"
	CorsDefaultAllowCredentials = "true"
	CorsDefaultMaxAge           = "86400"
)

type CorsOption struct {
	AllowOrigin      string
	AllowMethods     string
	AllowHeaders     string
	ExposeHeaders    string
	AllowCredentials string
	MaxAge           string
}

// Cors is a middleware function that appends headers to allow CORS with default config
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", CorsDefaultAllowOrigin)
		c.Header("Access-Control-Allow-Methods", CorsDefaultAllowMethods)
		c.Header("Access-Control-Allow-Headers", CorsDefaultAllowHeaders)
		c.Header("Access-Control-Expose-Headers", CorsDefaultExposeHeaders)
		c.Header("Access-Control-Allow-Credentials", CorsDefaultAllowCredentials)
		c.Header("Access-Control-Max-Age", CorsDefaultMaxAge)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}

}

// CorsWithOption is a middleware function that appends headers to allow CORS with option
func CorsWithOption(option CorsOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		if option.AllowOrigin != "" {
			c.Header("Access-Control-Allow-Origin", option.AllowOrigin)
		}
		if option.AllowMethods != "" {
			c.Header("Access-Control-Allow-Methods", option.AllowMethods)
		}
		if option.AllowHeaders != "" {
			c.Header("Access-Control-Allow-Headers", option.AllowHeaders)
		}
		if option.ExposeHeaders != "" {
			c.Header("Access-Control-Expose-Headers", option.ExposeHeaders)
		}
		if option.AllowCredentials != "" {
			c.Header("Access-Control-Allow-Credentials", option.AllowCredentials)
		}
		if option.MaxAge != "" {
			c.Header("Access-Control-Max-Age", option.MaxAge)
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
