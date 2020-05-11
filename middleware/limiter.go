package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/pkg/app"
	"golang.org/x/time/rate"
	"time"
)

func Limiter(maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second*1), maxBurstSize)
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
			return
		}
		fmt.Println("Too many requests")
		utilGin := app.Gin{C: c}
		utilGin.ResponseOk(errors.New("Too many requests"), nil)
		c.Abort()
		return
	}
}
