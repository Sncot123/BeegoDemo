package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {

	beego.Router("/orm1", &controllers.TestOrmController{})
}
