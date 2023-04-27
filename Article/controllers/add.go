package controllers

import (
	"Article/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AddController struct {
	beego.Controller
}

func (a *AddController) Get() {
	a.TplName = "add.html"
}
func (a *AddController) Post() {
	var article models.Article
	err := a.ParseForm(&article)
	if err != nil {
		a.Data["data"] = fmt.Sprintf("添加时表单解析错误!!!，%s", err)
		a.TplName = "error.html"
		return
	}
	o := orm.NewOrm()
	isCreate, i, err := o.ReadOrCreate(&article, "Title", "Author", "Content")
	if nil != err {
		a.Data["data"] = fmt.Sprintf("搜索或插入错误！！！,%s,%d", err, i)
		a.TplName = "error.html"
		return

	}
	if isCreate {
		a.Data["data"] = fmt.Sprintf("插入数据成功！！！")
		a.TplName = "success.html"
	}
}
