package orm

import (
	"fmt"
	"github.com/hero1s/ginweb/conf"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func getDataSourceName(username, passwd, host, schema string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&timeout=2s&readTimeout=5s&writeTimeout=5s&loc=Local", username, passwd, host, schema)
}

func InitDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine("mysql", getDataSourceName(c.UserName, c.Passwd, c.Host, c.Schema))
	if err != nil {
		panic(err)
	}
	db.ShowSQL(c.ShowSQL)
	db.SetMaxIdleConns(c.Idle)
	f, err := os.Create(c.LogFile)
	if err == nil {
		db.SetLogger(log.NewSimpleLogger(f))
		db.SetLogLevel(log.LogLevel(c.LogLevel))
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
