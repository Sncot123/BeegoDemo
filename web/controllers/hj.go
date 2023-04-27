package controllers

import "github.com/astaxie/beego"

type HjController struct {
	beego.Controller
}

func (h *HjController) Get() {
	h.TplName = "hj.html"
}
