package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID set X-Request-Id
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check for incoming header, use it if exist
		requestid := c.Request.Header.Get("X-Request-Id")
		// create request id with UUID4
		if requestid == "" {
			u4 := uuid.NewV4()
			requestid = u4.String()
		}
		// expose it for use in the application
		c.Set("X-Request-Id", requestid)
		c.Writer.Header().Set("X-Request-Id", requestid)
		c.Next()
	}
}
