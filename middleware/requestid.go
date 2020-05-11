package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/uuid"
)

func RequestId(c *gin.Context) {
	requestId := c.Request.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = uuid.GenUUID()
	}
	c.Set("X-Request-Id", requestId)
	c.Writer.Header().Set("X-Request-Id", requestId)
	c.Next()
}
