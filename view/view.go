package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTML render html
func HTML(ctx *gin.Context, name string, data any) {
	ctx.HTML(http.StatusOK, name, data)
}
