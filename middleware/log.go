package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/pkg/log"
	"time"
)

// 日志记录到文件
func LoggerToFile(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	// 处理请求
	c.Next()
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	// 请求方式
	reqMethod := c.Request.Method
	// 请求路由
	reqUri := c.Request.RequestURI
	// 状态码
	statusCode := c.Writer.Status()
	// 请求IP
	clientIP := c.ClientIP()
	// 日志格式
	log.Info("| %3d | %13v | %15s | %s | %s |",
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri,
	)
}

// 日志记录到 MongoDB
func LoggerToMongo(c *gin.Context) {

}

// 日志记录到 ES
func LoggerToES(c *gin.Context) {

}

// 日志记录到 MQ
func LoggerToMQ(c *gin.Context) {

}
