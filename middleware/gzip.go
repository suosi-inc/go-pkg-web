// Package middleware gzip to enable GZIP support.
// Ref: https://github.com/gin-contrib/gzip
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/suosi-inc/go-pkg-web/middleware/gzips"
)

func Gzip(level int, options ...gzips.Option) gin.HandlerFunc {
	return gzips.NewGzipHandler(level, options...).Handle
}
