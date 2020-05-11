module github.com/hero1s/ginweb

go 1.14

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gansidui/geohash v0.0.0-20141019080235-ebe5ba447f34
	github.com/gansidui/nearest v0.0.0-20141019122829-a5d0cde6ef14
	github.com/garyburd/redigo v1.6.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/howeyc/fsnotify v0.9.0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.10.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/spf13/viper v1.7.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/zheng-ji/goSnowFlake v0.0.0-20180906112711-fc763800eec9
	google.golang.org/appengine v1.6.6 // indirect
	gopkg.in/fatih/set.v0 v0.2.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/astaxie/beego v1.12.1 => github.com/nicle-lin/beego v1.12.7
