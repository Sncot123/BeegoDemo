package controllers

import (
	"Article/models"
	"Article/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateController struct {
	beego.Controller
}

func (u *UpdateController) Get() {
	iD := u.GetString("id")
	//fmt.Println(iD)
	id := utils.StrToInt(iD)
	fmt.Println(id)
	o := orm.NewOrm()
	article := models.Article{Id: id}
	//========================
	//fmt.Println(article)
	err := o.Read(&article)
	if nil != err {
		if err == orm.ErrNoRows {
			u.Data["data"] = fmt.Sprintf("数据库中没有此条数据，插入失败！！！%s", err)
			u.TplName = "error.html"
			return
		}
		u.Data["data"] = fmt.Sprintf("更新时遇到未知错误，%s", err)
		u.TplName = "error.html"
		return
	}
	u.Data["Article"] = article
	u.TplName = "update.html"
}

func (u *UpdateController) Post() {
	iD := u.GetString("id")
	id := utils.StrToInt(iD)
	var article models.Article
	err := u.ParseForm(&article)
	if nil != err {
		u.Data["data"] = fmt.Sprintf("更新操作表单解析失败，%s", err)
		u.TplName = "error.html"
		return
	}
	//========================
	article.Id = id
	fmt.Println(article)
	o := orm.NewOrm()
	_, err = o.Update(&article, "Title", "Author", "Content")
	if nil != err {
		u.Data["data"] = fmt.Sprintf("更新操作失败，%s", err)
		u.TplName = "error.html"
		return
	}
	u.Data["data"] = fmt.Sprintf("更新操作成功！！")
	u.TplName = "success.html"
}
