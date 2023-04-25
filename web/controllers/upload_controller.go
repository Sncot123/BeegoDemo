package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.TplName = "upload.html"
}

func (u *UploadController) Post() {
	file, header, err := u.GetFile("upload_file")
	if nil != err {
		return
	}
	defer func() {
		err2 := file.Close()
		if nil != err2 {
			return
		}
	}()
	filename := header.Filename
	fmt.Println(filename)
	unix := time.Now().Unix()
	formatInt := strconv.FormatInt(unix, 10)
	err = u.SaveToFile("upload_file", "static/upload/"+formatInt+"-"+filename)
	if nil != err {
		return
	}
	u.TplName = "success.html"
}
