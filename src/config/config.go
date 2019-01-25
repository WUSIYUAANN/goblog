package config

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "goblog/src/logs"
	"goblog/src/utils/bizerror"
	"goblog/src/utils/constant"
)

type dbConfig struct {
	DbUser         string
	DbPwd          string
	DbUrls         string
	DbName         string
	DbMaxIdleConns int
	DbMaxConns     int
	DbForce        bool
	DbDebug        bool
}

var (
	DB *dbConfig
)

func init() {
	err := beego.LoadAppConfig("ini", "conf/beego.conf")
	bizerror.Check(err)

	//日志配置
	logConf()
	//数据库信息
	dbConf()
	//函数导出信息
	funcConf()

	Log.Printf("goblog starup successful appPath:%v", beego.AppPath)
}

func dbConf() {
	cfg := beego.AppConfig
	DB = new(dbConfig)
	DB.DbUser = cfg.DefaultString("mysqlUser", "root")
	DB.DbPwd = cfg.DefaultString("mysqlPass", "lyd")
	DB.DbUrls = cfg.DefaultString("mysqlUrls", "localhost:3306")
	DB.DbName = cfg.DefaultString("mysqlDb", "goblog")
	DB.DbMaxIdleConns = cfg.DefaultInt("mysqlMaxIdleConns", 10)
	DB.DbMaxConns = cfg.DefaultInt("mysqlMaxOpenConns", 50)
	DB.DbForce = cfg.DefaultBool("mysqlForce", false)
	DB.DbDebug = cfg.DefaultBool("mysqlDebug", true)
	orm.Debug = true
	orm.NewLog(GetLogsWriter())

	Log.Printf("call Config init DB:%+v", DB)
}

func logConf() {
	InitLogs()
}

func funcConf() {
	cfg := beego.AppConfig
	err := beego.AddFuncMap("getValue", constant.GetValue)
	bizerror.Check(err)
	err = beego.AddFuncMap("appName", func() string { return cfg.String("appname") })
	bizerror.Check(err)
}
