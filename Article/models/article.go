package models

import "github.com/astaxie/beego/orm"

type Article struct {
	Id      int
	Title   string `form:"title"`
	Author  string `form:"author"`
	Content string `form:"content"`
}

func (a *Article) TableName() string {
	return "article"
}
func init() {
	orm.RegisterModel(new(Article))
}
