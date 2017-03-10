package middleware

import (
	"avalon/config"
	"avalon/util"
	"net/http"
	"strings"

	"gopkg.in/gin-gonic/gin.v1"
)

// RequestType check header required
func RequestType() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		contentType := headers.Get("Content-Type")

		if len(contentType) < 1 {
			c.JSON(http.StatusUnsupportedMediaType, util.FailResponse(config.ErrNotSupportedHeader.Error()))
			c.Abort()
		}

		if len(contentType) >= 1 && !strings.Contains(contentType, "application/json") {
			c.JSON(http.StatusUnsupportedMediaType, util.FailResponse(config.ErrNotSupportedHeader.Error()))
			c.Abort()
		}

		c.Next()
	}
}
