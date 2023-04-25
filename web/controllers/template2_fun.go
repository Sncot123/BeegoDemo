package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Template2Controller struct {
	beego.Controller
}

func (t *Template2Controller) Get() {

	t.Data["fun1"] = Test1
	t.Data["fun2"] = Test2
	t.Data["maps"] = map[string]string{"name": "潘丽萍"}
	t.Data["array"] = []int{1, 2, 3, 4}
	t.Data["str"] = "afasdfsd"
	t.TplName = "template2_fun.html"
}

func Test1() string {
	return "潘丽萍！"
}

func Test2(name string) string {
	return fmt.Sprintf("%s,我有点想你了", name)
}
