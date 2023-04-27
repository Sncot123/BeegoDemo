package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"web/models"
)

type TestOrmController struct {
	beego.Controller
}

func (t *TestOrmController) Get() {
	//user := models.Users{Name: "潘丽萍", Age: 19, Salary: 5000, Rate: 1}

	o := orm.NewOrm() //创建一个orm对象

	//o.Using("books")
	//insert, err := o.Insert(&user)
	//if nil != err {
	//	return
	//}
	//fmt.Println(insert)

	user1 := models.Users{Name: "潘丽萍"}
	o.Read(&user1, "Name")

	fmt.Println(user1)
	user2 := models.Users{Id: 8}
	o.Read(&user2)
	fmt.Println(user2)
	t.TplName = "v_orm1.html"
}
