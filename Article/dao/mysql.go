package dao

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	user := beego.AppConfig.String("user")
	pwd := beego.AppConfig.String("password")
	ip := beego.AppConfig.String("ip")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, dbname)
	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if nil != err {
		fmt.Println("数据库连接错误！")
		panic(err)
	}
}
