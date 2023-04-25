package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RouterController struct {
	beego.Controller
}

func (r *RouterController) Login() {
	param := r.Ctx.Input.Param(":id")
	fmt.Println(param, "--------------------")
	r.TplName = "router_get.html"
}
func (r *RouterController) Regist() {
	param := r.Ctx.Input.Param(":id")
	fmt.Println(param)
	r.TplName = "router_post.html"

}
