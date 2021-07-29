package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "productgo/routers"
	"productgo/utils"
)

func main() {
	if err := utils.InitMySQL(); err != nil {
		panic(err)
	}
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
