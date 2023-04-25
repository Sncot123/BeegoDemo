package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	beego.Router("/static", &controllers.StaticController{})
	beego.Router("/param", &controllers.ParamController{})
	beego.Router("/param/?:name", &controllers.ParamController{})
	beego.Router("/params", &controllers.ParamsControllers{})
	beego.Router("/paramsy", &controllers.TFormController{})
	beego.Router("/tajax", &controllers.TajaxController{})
	beego.Router("/flash", &controllers.FlashController{})
	//上传文件
	beego.Router("/upload", &controllers.UploadController{})
}
