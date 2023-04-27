package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Users))
}

type Users struct {
	Id     int    `orm:"auto"`
	Name   string `orm:"column(name)"`
	Age    int    `orm:"column(age)"`
	Salary int    `orm:"column(salary)"`
	Rate   int    `orm:"column(rate)"`
}

func (u *Users) TableName() string {
	return "users"
}
