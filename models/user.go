package models

import (
	"fmt"
	"productgo/utils"
)

type User struct {
	Id         int
	Username   string
	Password   string
	CreateTime int64
}

func InsertUser(user *User) (int64, error) {
	QueryUserWithUserName(user.Username)
	sql := "insert into user(username,password,createtime) values(?,?,?)"
	return utils.ExecSQL(sql, user.Username, user.Password, user.CreateTime)
}

func QueryUserWithUserName(username string) int {
	sql := fmt.Sprintf("select id from user where username=%s", username)
	row:=utils.QueryRow(sql)
	id:=0
	row.Scan(&id)
	return  id
}
