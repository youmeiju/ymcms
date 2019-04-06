package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"time"
)

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController)DelGoods()  {
	id:=this.GetString("goodsId")
	beego.Error("id=",id)
	o:=orm.NewOrm()
	o.Raw("delete from goods_info where id=?",id).Exec()
	this.Redirect("/index",302)
}

func (this *GoodsController)ShowAddGoods()  {
	msg:=this.GetString("msg")
	o:=orm.NewOrm()
	var types []orm.Params
	o.Raw("SELECT * from first_menu").Values(&types)
	this.Data["types"]=types
	this.Data["msg"]=msg
	this.TplName = "add.html"
}


func (this *GoodsController)HandleAddGoods()  {
	goodsPrice,err:=this.GetInt("goodsPrice")
	if err!=nil{
		beego.Error("获取商品价格失败",err)
		this.Redirect("/addGoods?msg=获取商品价格失败",302)
		return
	}
	goodsName:=this.GetString("goodsName")
	if goodsName==""{
		beego.Error("获取商品名称失败")
		this.Redirect("/addGoods?msg=获取商品名称失败",302)
		return
	}
	goodsImg:=UpLoad1(this,"goodsImg")
	goodsPhoto:=UpLoad1(this,"goodsPhoto")
	o:=orm.NewOrm()
	typeName:=this.GetString("select")
	var maps []orm.Params
	o.Raw("select * from first_menu where name=?",typeName).Values(&maps)
	o.Raw("INSERT INTO goods_info VALUES(null,?,?,?,?,?,0)",goodsName,maps[0]["id"],goodsImg,goodsPhoto,goodsPrice).Exec()
	this.Redirect("/addGoods?msg=添加商品成功",302)
}

func UpLoad1(this *GoodsController,filePath string)(string)  {
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

func (this *GoodsController)ShowAddGoodsType()  {
	err:=this.GetString("err")
	if err!=""{
		this.Data["Err"]=err
	}
	o:=orm.NewOrm()
	var Types []orm.Params
	o.Raw("SELECT * FROM first_menu").Values(&Types)
	this.Data["Types"]=Types
	this.TplName = "addType.html"
}

func (this *GoodsController)HandleAddType()  {
	name:=this.GetString("name")
	if name==""{
		this.Redirect("/addType?err=请填写商品分类名称",302)
	}
	o:=orm.NewOrm()
	o.Raw("insert into first_menu values(null,?)",name).Exec()
	this.Redirect("/addType",302)
}

func (this *GoodsController)DelTypes()  {
	id:=this.GetString("id")
	o:=orm.NewOrm()
	var lists []orm.ParamsList
	o.Raw("select count(id) from goods_info where g_mf_id=3;").ValuesList(&lists)
	if lists[0][0]!=0{
		this.Redirect("/addType?err="+"请先删除当前分类下的商品才能删除该分类",302)
		return
	}
	o.Raw("DELETE FROM first_menu where id=?;",id).Exec()
	this.Redirect("/addType",302)
}

