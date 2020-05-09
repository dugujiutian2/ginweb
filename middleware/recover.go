package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/pkg/log"
	"net/http"
	"net/http/httputil"
	"runtime"
	"time"
)

func RecoverHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			httprequest, _ := httputil.DumpRequest(c.Request, false)
			pnc := fmt.Sprintf("[Recovery] %s panic recovered:\n%s\n%s\n%s", time.Now().Format("2006-01-02 15:04:05"), string(httprequest), err, buf)
			fmt.Print(pnc)
			log.Error(pnc)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}
