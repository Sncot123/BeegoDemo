package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Router("/template2", &controllers.Template2Controller{})
}
