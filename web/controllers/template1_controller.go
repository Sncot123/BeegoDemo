package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
)

type Template1Controller struct {
	beego.Controller
}

func (t *Template1Controller) Get() {

	t.Data["name"] = "你好！潘丽萍"
	t.Data["arr"] = []int{1, 2, 3, 4, 5}
	t.Data["vip"] = 1
	t.Data["stuc"] = models.User{UserName: "潘丽萍", Age: 19}

	t.TplName = "template1_grammer.html"
}
