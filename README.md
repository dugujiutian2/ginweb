# ginweb
基于gin的web组件工具集  
1:db orm        
2:cache     
3:log   
4:鉴权    
5:路由过滤授权，注册封装   
6:异常处理  
7:https 
8:session   
9:性能监控  
10:在线文档  
11:跨域处理  
12:task  
13:config      
14:RPC  

安装swag命令    
go get github.com/swaggo/swag/cmd/swag  
项目根目录里执行swag init，生成docs/docs.go；访问http://localhost:8080/swagger/index.html 

#### 运行参数：

-env 可选，根据配置文件而定
* prod
* test

```bash
-env prod -conf cmd/web_template.json
```

#### 数据库 Reverse
工具自动生成方法:使用xorm工具 
安装工具:go get github.com/go-xorm/cmd/xorm 
```bash
xorm reverse mysql dbUserName:password@tcp(hostname:3306)/databaseName?charset=utf8 dbmodel/goxorm dbmodel
```