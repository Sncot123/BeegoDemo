package controllers

import (
	"Article/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TestController struct {
	beego.Controller
}

func (t *TestController) Get() {
	o := orm.NewOrm()
	//使用表名
	//qs := o.QueryTable("article")

	//使用结构体
	model := models.Article{}
	qs := o.QueryTable(model)
	articles := []models.Article{}
	orm.Debug = true //开启debug模式  仅在当前controller模块下  在控制台打印原生sql语句
	//[SELECT T0.`id`, T0.`title`, T0.`author`, T0.`content` FROM `article` T0 WHERE T0.`author` LIKE BINARY ? AND NOT T0.`content` LIKE ? LIMIT 3 OFFSET 1] - `%li%`, `%LOVE%`
	_, err := qs.Filter("author__contains", "li").Limit(4).Offset(1).All(&articles)
	if nil != err {
		t.Data["data"] = fmt.Sprintf("过滤查询失败！！！%s", err)
		t.TplName = "error.html"
		return
	}
	t.Data["articles"] = articles
	t.TplName = "test.html"
	//https: //subs.mm2023.top//link/srFq8nbKwk2czR20?sub=4
}
