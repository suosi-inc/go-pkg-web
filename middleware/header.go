package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suosi-inc/go-webpkg/rest"
)

// HeaderAuth is a middleware function that validates the request header.
// Return http forbidden code or rest json
func HeaderAuth(k string, v string, json bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get(k) != v {
			if json {
				rest.JsonError(c, http.StatusForbidden)
				c.Abort()
			} else {
				c.AbortWithStatus(http.StatusForbidden)
			}
		} else {
			c.Next()
		}
	}
}
