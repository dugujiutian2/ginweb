package orm

import (
	"github.com/hero1s/ginweb/conf"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func InitDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Driver, c.Dsn)
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
