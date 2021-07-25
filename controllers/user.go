package controllers

import beego "github.com/beego/beego/v2/server/web"

type LoginController struct {
	beego.Controller
}

type  RegisterController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}


func (r *RegisterController) Get() {
	r.TplName = "register.html"
}
