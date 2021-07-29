package models


type Product struct{
	Id int
	Name string
	Link string
	Description  string
	Author  User
	CreateTime  int64
	Enable  bool
}
