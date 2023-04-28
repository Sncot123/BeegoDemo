package routers

import (
	"Article/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/update/?:id", &controllers.UpdateController{})
	beego.Router("/delete/?:id", &controllers.DeleteController{})
	beego.Router("/test", &controllers.TestController{})

}
