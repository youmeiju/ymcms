package main

import (
	_ "ymcms/routers"
	"github.com/astaxie/beego"
	_"ymcms/models"
)

func main() {
	beego.AddFuncMap("prev",getPrev)
	beego.AddFuncMap("next",getNext)
	beego.Run()
}

func getPrev(pageId int)int  {
	if pageId-1<1{
		return 1
	}
	return pageId-1
}

func getNext(pageId int,pageCount int)int  {
	if pageId+1>pageCount{
		return pageCount
	}
	return pageId+1
}