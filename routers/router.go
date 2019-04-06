package routers

import (
	"ymcms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/index",&controllers.IndexController{},"get:ShowIndex")
	beego.Router("/update",&controllers.UpdateController{},"get:ShowUpdate;post:HandleUpdate")
	beego.Router("/addType",&controllers.GoodsController{},"get:ShowAddGoodsType;post:HandleAddType")
	beego.Router("/delType",&controllers.GoodsController{},"get:DelTypes")
	beego.Router("/deleteGoods",&controllers.GoodsController{},"get:DelGoods")
	beego.Router("/addGoods",&controllers.GoodsController{},"get:ShowAddGoods;post:HandleAddGoods")
}
