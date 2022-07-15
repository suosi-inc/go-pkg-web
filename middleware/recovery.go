package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suosi-inc/go-webpkg/rest"
)

// RestRecovery For gin.CustomRecovery
func RestRecovery(stack bool) gin.RecoveryFunc {
	return func(c *gin.Context, err interface{}) {
		if stack {
			e := fmt.Sprintf("%v", err)
			rest.JsonErrorMessage(c, http.StatusInternalServerError, e)
		} else {
			rest.JsonError(c, http.StatusInternalServerError)
		}
		c.Abort()
	}
}
