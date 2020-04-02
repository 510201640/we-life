package initialize

import (
	"common-lib/mysql"
	"jaden/we-life/common"
)

func SetupDB() {
	setUpDBMaster()

}

func setUpDBMaster() {
	common.DB = mysql.New(mysql.Config{
		Dialect:  common.Cfg.GetString("mysql.db.dialect"),
		User:     common.Cfg.GetString("mysql.db.user"),
		Password: common.Cfg.GetString("mysql.db.password"),
		DataBase: common.Cfg.GetString("mysql.db.database"),
		Host:     common.Cfg.GetString("mysql.db.host"),
		Port:     common.Cfg.GetString("mysql.db.port"),
		MaxIdle:  common.Cfg.GetInt("mysql.db.maxIdle"),
		MaxOpen:  common.Cfg.GetInt("mysql.db.maxOpen"),
	})

}
