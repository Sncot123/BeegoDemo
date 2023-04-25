package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	//固定路由
	//beego.Router("/router", &controllers.RouterController{})
	//正则路由
	//beego.Router("/router/:id[0-9]+", &controllers.RouterController{})
	////自动路由
	//beego.AutoRouter(&controllers.RouterController{})
	////自定义路由
	beego.Router("/router", &controllers.RouterController{}, "get:Login;post:Regist")

}
