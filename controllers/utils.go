package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"time"
)

func UpLoad(this *beego.Controller,filePath string)string  {
	file,head,err :=this.GetFile(filePath)
	//校验数据
	if err != nil{
		beego.Error(err)
		return ""
	}
	defer file.Close()
	//1.文件存在覆盖的问题
	//加密算法

	//当前时间
	fileName := time.Now().Format("2006-01-02-15-04-05")
	ext := path.Ext(head.Filename)
	beego.Info(head.Filename,ext)
	//2.文件类型也需要校验
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		beego.Error(err)
		return ""
	}
	//3.文件大小校验
	if head.Size > 5000000 {
		beego.Error(err)
		return ""
	}

	//把图片存起来
	err=this.SaveToFile(filePath,"/root/go/src/NewService/img/goods/"+fileName+ext)
	if err!=nil{
		beego.Error(err)
		return ""
	}
	return "https://service.shanghaiyoumeiju2018.com/img/goods/"+fileName+ext
}