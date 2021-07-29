package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//设置模板
	c.TplName = "index.html"
	//uid := c.GetSession("login_uid")
	//if uid != nil {
	//	c.Data["Uid"] = uid
	//}
}
