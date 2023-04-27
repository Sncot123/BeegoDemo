package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "web/routers"
)

func init() {
	//err := orm.RegisterDriver("mysql", orm.DRMySQL)
	//if nil != err {
	//	return
	//}
	err := orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/books?charset=utf8")
	if nil != err {
		return
	}

}

func main() {
	//beego.SetViewsPath("")

	//s := beego.AppConfig.String("runmode") ///获取配置文件的值bee
	//
	//fmt.Println(s)
	//beego.SetStaticPath("/static", "front")
	//
	//beego.BConfig.MaxMemory = 1 << 22 //单位时B
	//beego.BConfig.Listen.EnableAdmin = true
	beego.Run()
}
