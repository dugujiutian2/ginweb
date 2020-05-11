package orm

import (
	"github.com/hero1s/ginweb/conf"
	"xorm.io/xorm"
)

func InitDB(c *conf.DB) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Driver, c.Dsn)
	if err != nil {
		panic(err)
	}
	db.ShowSQL(c.ShowSQL)
	db.SetMaxIdleConns(c.Idle)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
