package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
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
	goodsImg:=UpLoad(&this.Controller,"goodsImg")
	goodsPhoto:=UpLoad(&this.Controller,"goodsPhoto")
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