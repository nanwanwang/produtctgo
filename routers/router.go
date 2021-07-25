package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"productgo/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/login",&controllers.LoginController{})
}
