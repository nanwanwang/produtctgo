package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"productgo/models"
	"productgo/utils"
	"time"
)

type LoginController struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	fmt.Println(username, password, repassword)
	if password != repassword {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "两次密码不一样",
		}
		r.ServeJSON()
		return
	}

	id := models.QueryUserWithUserName(username)
	if id > 0 {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "用户名已存在",
		}
		r.ServeJSON()
		return
	}
	user := models.User{0, username, utils.MD5(password), time.Now().Unix()}
	_, err := models.InsertUser(&user)
	if err != nil {
		r.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "注册失败",
		}
	} else {
		r.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "注册成功",
		}
	}
	r.ServeJSON()

}


func (l *LoginController) Post(){
	username:=l.GetString("username")
	password:=l.GetString("password")
	fmt.Println(username,password)

	id := models.QueryUserWithParam(username,utils.MD5(password))
	if id <= 0{
		l.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "登录失败",
		}
	}else{
		l.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "登录成功",
		}
		l.SetSession("login_uid",id)
	}
	l.ServeJSON()
}
