package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const XRequestIDKey = "XRequestID"

// RequestIDMiddleware 生成唯一请求 ID 并将其存储在上下文中
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(XRequestIDKey, requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
