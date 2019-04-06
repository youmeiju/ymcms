package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"ymcms/models"
	"regexp"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)ShowLogin()  {
	this.TplName = "login.html"
}

func (this *UserController)HandleLogin()  {
	userName:=this.GetString("userName")
	passWord:=this.GetString("password")
	//校验用户名和密码
	if userName==""||passWord==""{
		beego.Error("账号或密码不能为空")
		this.Data["Err"] = "账号或密码不能为空"
		this.TplName = "login.html"
		return
	}
	reg:=regexp.MustCompile(`^1\d{10}$`)
	b:=reg.MatchString("13122778585")
	if !b{
		beego.Error("手机格式号格式不正确")
		this.Data["Err"] = "手机格式号格式不正确"
		this.TplName = "login.html"
		return
	}
	//orm操作
	o:=orm.NewOrm()
	user:=&models.AdminUser{
		TelNum:userName,
	}
	err:=o.Read(user,"TelNum")
	if err!=nil{
		beego.Error("账号不存在",err)
		this.TplName = "login.html"
		return
	}
	if user.PassWord!=passWord{
		beego.Error( "密码错误，请修正后重试")
		this.Data["Err"] = "密码错误，请修正后重试"
		this.TplName = "login.html"
		return
	}
	this.Redirect("/index",302)
}