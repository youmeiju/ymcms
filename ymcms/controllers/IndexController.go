package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"time"
	"ymcms/models"
)

type IndexController struct {
	beego.Controller
}
//展示首页
func (this *IndexController)ShowIndex()  {
	pageId,err:=this.GetInt("pageId")
	if err!=nil{
		pageId=1
	}
	o:=orm.NewOrm()
	qs:=o.QueryTable("GoodsInfo")
	//var goodsInfos []models.GoodsInfo
	goodsCount,err:=qs.Count()
	if err!=nil{
		beego.Error("获取数量失败",err)
		this.Data["Err"] = "获取商品列表失败，请稍后重试"
		this.TplName = "index.html"
		return
	}
	page:=15
	pageCount:=math.Ceil(float64(goodsCount)/float64(page))
	//获取并设置分页信息
	this.Data["pageCount"] = int(pageCount)
	this.Data["goodsCount"] = goodsCount
	this.Data["page"] = page
	this.Data["pageId"]=pageId
	//获取商品信息
	start:=int(pageId-1)*page
	var goodsInfos []models.GoodsInfo
	_,err=qs.Limit(page,start).All(&goodsInfos)
	if err!=nil{
		beego.Error("获取数量失败",err)
		this.Data["Err"] = "获取商品列表失败，请稍后重试"
		this.TplName = "index.html"
		return
	}
	this.Data["Goods"] = goodsInfos
	this.TplName = "index.html"
}


type UpdateController struct {
	beego.Controller
}

//编辑商品信息
func (this *UpdateController)ShowUpdate()  {
	id,err:=this.GetInt("goodsId")
	if err!=nil{
		beego.Error("获取商品id失败",err)
		this.Redirect("/index",302)
		return
	}
	o:=orm.NewOrm()
	var goodsInfo models.GoodsInfo
	goodsInfo.Id = id
	err=o.Read(&goodsInfo)
	if err!=nil{
		beego.Error("获取商品信息失败",err)
		this.Redirect("/index",302)
		return
	}
	this.Data["goodsInfo"] = goodsInfo
	this.TplName = "update.html"
}

//更改商品信息
func (this *UpdateController)HandleUpdate()  {
	id,err:=this.GetInt("Id")
	if err!=nil{
		beego.Error("获取商品id失败",err)
		this.Redirect("/index",302)
		return
	}
	goodsPrice,err:=this.GetInt("goodsPrice")
	if err!=nil{
		beego.Error("获取商品价格失败",err)
		this.Redirect("/index",302)
		return
	}
	goodsName:=this.GetString("goodsName")
	if goodsName==""{
		beego.Error("获取商品名称失败")
		this.Redirect("/index",302)
		return
	}
	goodsImg:=UpLoad(this,"goodsImg")
	goodsPhoto:=UpLoad(this,"goodsPhoto")
	o:=orm.NewOrm()
	var goodsInfo models.GoodsInfo
	goodsInfo.Id = id
	err=o.Read(&goodsInfo)
	if err!=nil{
		beego.Error("商品不存在")
		this.Redirect("/index",302)
		return
	}
	if goodsImg!=""{
		goodsInfo.GImg = goodsImg
	}
	if goodsPhoto!=""{
		goodsInfo.GPhoto = goodsPhoto
	}
	goodsInfo.ShopPrice = goodsPrice
	goodsInfo.GName = goodsName
	_,err=o.Update(&goodsInfo)
	if err!=nil{
		beego.Error("更新商品信息失败")
		this.Redirect("/index",302)
		return
	}
	this.Redirect("/index",302)
}

func saveImg(this *UpdateController,fliePath string)(string,error)  {
	file,head,err:=this.GetFile(fliePath)
	if err!=nil{
		beego.Error("获取图片信息失败",err)
		return "",err
	}
	defer file.Close()
	ext:=path.Ext(head.Filename)
	//2.文件类型也需要校验
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg"{
		beego.Error("上传图片格式不正确，请重新上传",head.Filename)
		return "",err
	}
	//3.文件大小校验
	if head.Size > 50000000 {
		beego.Error("上传图片格式不正确，请重新上传",head.Size)
		return "",err
	}
	upTime:=time.Now().Format("2006-01-02 15:04:05")
	fileName := upTime+ext
	this.SaveToFile(fliePath,"/home/wujiu/Desktop/"+fileName)
	return "/home/wujiu/Desktop/"+fileName,err
}

func UpLoad(this *UpdateController,filePath string)(string)  {
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