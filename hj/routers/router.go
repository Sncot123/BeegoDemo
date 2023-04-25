package routers

import (
	"github.com/astaxie/beego"
	"hj/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
