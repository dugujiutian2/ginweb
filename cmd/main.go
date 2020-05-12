package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/conf"
	_ "github.com/hero1s/ginweb/docs"
	"github.com/hero1s/ginweb/middleware"
	"github.com/hero1s/ginweb/pkg/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title gin服务框架测试
// @version 0.0.1
// @description  测试文档
// @BasePath /api/v1/
func main() {
	log.InitLog("dev", "logs")
	// ParseConfig
	err := conf.ParseConfig()
	if err != nil {
		log.Error("解析配置文件错误:%+v",err)
	}

/*	db.InitDB("","","","","",true,0,log.DefaultLog,200,200)
	// init Db
	db := orm.InitDB(conf.Conf.DB)
	defer db.Close()*/

	r := gin.New()
	//开启中间件记录日志
	r.Use(middleware.LoggerToFile)
	r.Use(middleware.RecoverHandler)
	r.Use(middleware.Cors)
	//r.Use(middleware.JWT)

	// 创建路由组
	v1 := r.Group("/api/v1")

	v1.GET("/record/:userId", record)
	v1.GET("/record2/:userId", record2)
	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(conf.Conf.HttpServer.Port)
}

// @Summary 测试中文注解2
// @Description 根据用户ID获取信息2
// @Tags 测试分组2
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "用户ID"
// @Success 200 {string} string	"ok"
// @Router /record2/{some_id} [get]
func record2(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// @Summary 测试中文注解
// @Description 根据用户ID获取信息
// @Tags 测试分组
// @Accept  json
// @Produce json
// @Param   some_id     path    int     true        "用户ID"
// @Success 200 {string} string	"ok"
// @Router /record/{some_id} [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

