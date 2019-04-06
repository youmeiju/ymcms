package routers

import (
	"ymcms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/index",&controllers.IndexController{},"get:ShowIndex")
	beego.Router("/update",&controllers.UpdateController{},"get:ShowUpdate;post:HandleUpdate")
}
