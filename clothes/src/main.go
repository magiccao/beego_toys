package main

import (
	_ "models"
	"os"
	_ "routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// init orm
	dbDriver := beego.AppConfig.String("db_driver")
	dbSource := beego.AppConfig.String("db_source")
	dbTimeout, err := beego.AppConfig.Int("db_timeout")
	if err != nil {
		beego.Warn("db_timeout invalid. Use default value: 30")
		dbTimeout = 30
	}

	err = orm.RegisterDataBase("default", dbDriver, dbSource, dbTimeout)
	if err != nil {
		beego.Critical("orm.RegisterDataBase err: ", err.Error())
		beego.BeeLogger.Close()
		os.Exit(1)
	}
	orm.RunCommand()

	err = orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Critical("orm.RunSyncdb err: ", err.Error())
		beego.BeeLogger.Close()
		os.Exit(1)
	}

	beego.Run()
}
