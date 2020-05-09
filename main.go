package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/hero1s/ginweb/docs"
	"github.com/hero1s/ginweb/middleware"
	"github.com/hero1s/ginweb/pkg/db"
	"github.com/hero1s/ginweb/pkg/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title 测试gin服务
// @version 0.0.1
// @description  测试
// @BasePath /api/v1/
func main() {

	log.InitLog("dev", "logs")
	db.InitDB("","","","","",true,0,log.DefaultLog,200,200)

	r := gin.New()
	//开启中间件记录日志
	r.Use(middleware.LoggerToFile)
	r.Use(middleware.RecoverHandler)
	r.Use(middleware.Cors)
	r.Use(middleware.JWT)

	// 创建路由组
	v1 := r.Group("/api/v1")

	v1.GET("/record/:userId", record)

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @title 测试
// @获取指定ID记录
// @Description get record by ID
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "userId"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
