package controllers

import (
	"Article/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	articles := []models.Article{}
	o := orm.NewOrm()
	qt := o.QueryTable("article")
	_, err := qt.All(&articles)
	if err != nil {
		if err == orm.ErrNoRows {
			c.Data["data"] = fmt.Sprintf("未查到数据！！！%s", err)
			c.TplName = "error.html"
			return
		}
		c.Data["data"] = fmt.Sprintf("未知错误！！！%s", err)
		c.TplName = "error.html"
		return
	}
	//=======================
	//fmt.Println(article)
	c.Data["Articles"] = articles
	c.TplName = "index.html"
}
