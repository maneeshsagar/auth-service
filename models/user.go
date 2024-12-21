package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	ID       int64  `orm:"auto;column(id)"`
	Name     string `orm:"column(name)"`
	Email    string `orm:"column(email)"`
	Password string `orm:"column(password)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "user"
}
