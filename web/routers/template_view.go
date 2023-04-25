package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Router("/template1", &controllers.Template1Controller{})
}
