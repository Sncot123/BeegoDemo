package logs

import "github.com/astaxie/beego/logs"

func init() {
	logs.SetLogger("file", `{"filename":"logs/test_log.log"}`)
	logs.SetLevel(logs.LevelInformational)
}
