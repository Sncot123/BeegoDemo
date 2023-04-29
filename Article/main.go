package main

import (
	_ "Article/dao"
	_ "Article/routers"
	"github.com/astaxie/beego"
)

func main() {

	//beego.SetLogger("file", `{"filename":"logs/test_log.log"}`)
	//beego.BeeLogger.DelLogger("console") //删除控制台的日志输出
	//beego.SetLevel(beego.LevelError)     //设置日志级别
	//beego.SetLogFuncCall(true)           //设置日志输出无行号

	beego.Run()
}
