package dao

import (
	"fmt"
	"github.com/astaxie/beego/cache"
)

//  缓存的使用
var MemCache cache.Cache

func init() {
	adapter, err := cache.NewCache("memory", `{"interval":60}`)
	if nil != err {
		panic("缓存创建失败！！！！！")
	}
	err = adapter.Put("潘丽萍", "love you so much", 60)
	if nil != err {
		fmt.Errorf("数据存入失败！！！err:%s", err)
		return
	}
	MemCache = adapter
}
