package routers

import (
	"github.com/astaxie/beego"
	"vb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
