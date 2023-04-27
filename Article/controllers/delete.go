package controllers

import (
	"Article/models"
	"Article/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DeleteController struct {
	beego.Controller
}

func (d *DeleteController) Get() {
	iD := d.GetString("id")
	id := utils.StrToInt(iD)
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)
	if err != nil {
		d.Data["data"] = fmt.Sprintf("出现了错误！！！%s", err)
		d.TplName = "error.html"
		return
	}
	i, err := o.Delete(&article)
	if nil != err {
		d.Data["data"] = "删除操作出现了错误" + "-->" + string(i)
		d.TplName = "error.html"
		return
	}
	d.Data["data"] = "删除操作成功"
	d.TplName = "success.html"

}
