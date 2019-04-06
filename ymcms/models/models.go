package models

import "github.com/astaxie/beego/orm"
import _"github.com/go-sql-driver/mysql"

func init()  {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(118.31.72.110:3306)/YM?charset=utf8")
	orm.RegisterModel(new(AdminUser),new(GoodsInfo))
}


type AdminUser struct {
	Id int
	TelNum string `orm:"size(50)"`
	PassWord string `orm:"size(50)"`
	Openid string `orm:"size(50)"`
	IsSurper bool
	Token string `orm:"size(50)"`
	RealName string `orm:"size(50)"`
	PushOpenid string `orm:"size(50)"`
}

type GoodsInfo struct {
	Id int
	GName string `orm:"size(50)"`
	GMfId int
	GImg string `orm:"size(100)"`
	GPhoto string `orm:"size(100)"`
	ShopPrice int
	SaleValum int
}