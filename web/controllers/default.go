package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"web/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["Name"] = "潘丽萍"
	u.Data["Age"] = "18"
	arrays := [3]int{1, 2, 3}
	u.Data["arrays"] = arrays
	arrayStruct := [3]models.User{{"潘丽萍", 18}, {"于欢", 21}, {"樊雪姨", 23}}
	u.Data["array_struct"] = arrayStruct
	maps := map[string]string{"name": "潘丽萍", "age": "18", "address": "贵州"}
	u.Data["maps"] = maps
	mapStruct := map[int]models.User{1: {"潘丽萍", 1}, 2: {"樊雪姨", 23}, 3: {"于欢", 23}}
	u.Data["map_struct"] = mapStruct

	u.TplName = "user.html"
}

type StaticController struct {
	beego.Controller
}

func (s *StaticController) Get() {

	s.TplName = "sta.html"
}

type ParamController struct {
	beego.Controller
}

func (p *ParamController) Get() {
	//方式一   http://127.0.0.1:8080/param?name=12121
	name1 := p.GetString("name")
	fmt.Println("name1->", name1)

	name2 := p.Input().Get("name")
	fmt.Println("name2->", name2)
	//方式二	http://127.0.0.1:8080/param/123132
	//				/?:name
	name3 := p.Ctx.Input.Param(":name")
	fmt.Println("name3->", name3)
	name4 := p.GetString(":name")
	fmt.Println("name4->", name4)
	p.TplName = "para.html"

}

type ParamsControllers struct {
	beego.Controller
}

func (p *ParamsControllers) Get() {
	p.TplName = "tform.html"
}
func (p *ParamsControllers) Post() {

	name := p.GetString("username")
	age, _ := p.GetInt("userage")
	price, _ := p.GetFloat("price")
	single, _ := p.GetBool("single")
	p.Data["name"] = name
	p.Data["age"] = age
	p.Data["price"] = price
	p.Data["single"] = single
	fmt.Println(name, "->", age, "->", price, "->", single)
	p.TplName = "success.html"
}

type TFormController struct {
	beego.Controller
}

func (t *TFormController) Get() {
	t.TplName = "tform.html"

}
func (t *TFormController) Post() {
	var user models.TForm
	err := t.ParseForm(&user)
	if nil != err {
		return
	}
	fmt.Println(user)
	t.TplName = "success.html"
}

/**

测试ajax方式的数据传输
view层  tajax.html


*/

type TajaxController struct {
	beego.Controller
}

func (t *TajaxController) Get() {
	t.TplName = "tajax.html"
}
func (t *TajaxController) Post() {
	//获取ajax数据

	fmt.Println("-0-----------------------------------")
	var tajax models.Tajax
	body := t.Ctx.Input.RequestBody //二进制json数据
	_ = json.Unmarshal(body, &tajax)
	fmt.Println(tajax)

	result := map[string]string{"code": "200", "msg": "处理成功！！"}

	t.Data["json"] = result
	t.ServeJSON() //返回json格式
}

/**
  flash 数据
*/

type FlashController struct {
	beego.Controller
}

func (f *FlashController) Get() {
	flash := beego.ReadFromRequest(&f.Controller)
	notice := flash.Data["notice"]
	err := flash.Data["error"]
	warning := flash.Data["warning"]
	if len(warning) != 0 {
		f.TplName = "warningflash.html"
	} else if len(err) != 0 {
		f.TplName = "errorflash.html"
	} else if len(notice) != 0 {
		f.TplName = "noticeflash.html"
	} else {
		f.TplName = "flash.html"
	}

}

func (f *FlashController) Post() {

	//创建一个flash对象
	flash := beego.NewFlash()
	name := f.GetString("username")
	pwd := f.GetString("pwd")
	if name == "" {
		fmt.Println("用户名不能为空")
		flash.Warning("用户名不能为空！！！")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302) //重定向
	} else if pwd != "123456" {
		fmt.Println("密码错误")
		flash.Error("密码错误！！！")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302) //重定向
	} else {
		fmt.Println("登陆成功")
		flash.Notice("登陆成功！！！")
		flash.Store(&f.Controller)
		f.Redirect("/flash", 302) //重定向
	}

}
