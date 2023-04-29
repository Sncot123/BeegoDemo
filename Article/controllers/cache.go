package controllers

import (
	"Article/dao"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type CacheController struct {
	beego.Controller
}

func (c *CacheController) Get() {
	key := "潘丽萍"
	value := "潘丽萍，曾经我很喜欢你"
	err := dao.MemCache.Put(key, value, 180)
	if nil != err {
		c.Data["data"] = fmt.Sprintf("存储失败！！！err:%s", err)
		c.TplName = "error.html"
		return
	}
	logs.Emergency("存储成功")
	get := dao.MemCache.Get(key)
	logs.Emergency("取值成功，值：", get)
	c.TplName = "test_cache.html"
}
