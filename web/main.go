package main

import (
	"github.com/astaxie/beego"
	_ "web/routers"
)

func main() {
	//beego.SetViewsPath("")

	//s := beego.AppConfig.String("runmode") ///获取配置文件的值bee
	//
	//fmt.Println(s)
	//beego.SetStaticPath("/static", "front")
	//
	//beego.BConfig.MaxMemory = 1 << 22 //单位时B
	beego.Run()
}
