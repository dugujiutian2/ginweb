package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/conf"
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
	// ParseConfig
	err := conf.ParseConfig()
	if err != nil {
		log.Error("解析配置文件错误:%+v",err)
	}

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

	r.Run(conf.Conf.HttpServer.Port)
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

/*func studentById(c *gin.Context) error {
	pId := c.Query("id")
	if pId == "" {
		return (c, nil, ecode.RequestErr)
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}

	hystrix.Do(config.ProjectName, func() error {
		rsp, err := schoolSrv.StudentById(id)
		if err == nil {
			JSON(c, rsp, nil)
		}
		return err
	}, func(e error) error {
		switch e {
		case hystrix.ErrCircuitOpen:
			fmt.Println("studentList 熔断开启:", e)
			JSON(c, nil, ecode.New(502, "熔断开启"))
		case hystrix.ErrMaxConcurrency:
			fmt.Println("studentList 熔断超过最大并发:", e)
			JSON(c, nil, ecode.New(502, "熔断超过最大并发"))
		case hystrix.ErrTimeout:
			fmt.Println("studentList 熔断超时:", e)
			JSON(c, nil, ecode.New(502, "熔断超时"))
		default:
			JSON(c, nil, e)
		}
		return e
	})
	return nil
}*/

