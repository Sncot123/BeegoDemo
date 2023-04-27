package main

import (
	_ "Article/dao"
	_ "Article/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
