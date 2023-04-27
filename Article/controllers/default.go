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
	article := models.Article{Id: 3}
	o := orm.NewOrm()
	err := o.Read(&article)
	if err != nil {
		if err == orm.ErrNoRows {
			fmt.Println("未查到数据！！！！")
			c.TplName = "index.html"
			return
		}
		fmt.Println("未知错误！ err:", err)
		return
	}
	//=======================
	fmt.Println(article)
	c.Data["Article"] = article
	c.TplName = "index.html"
}
