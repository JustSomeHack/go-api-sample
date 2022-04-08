package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-Not-Valid") != "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "not allowed",
			})
			return
		}
		c.Next()
	}
}
