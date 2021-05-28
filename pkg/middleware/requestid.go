package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/constant"
)

// RequestID 透传Request-ID，如果没有则生成一个
func RequestID(c *gin.Context) {
	// Check for incoming header, use it if exists
	requestID := c.Request.Header.Get(constant.XRequestID)

	if requestID == "" {
		requestID = xid.New().String()
	}

	// Expose it for use in the application
	c.Set(constant.XRequestID, requestID)

	// Set X-Request-ID header
	c.Writer.Header().Set(constant.XRequestID, requestID)
}
