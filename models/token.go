package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Token struct {
	ID           int64     `orm:"column(id);auto"`
	UserID       int64     `orm:"column(user_id)"`
	RefreshToken string    `orm:"column(refresh_token)"`
	Token        string    `orm:"column(token)"`
	ExpiresAt    time.Time `orm:"column(expires_at)"`
	CreatedAt    time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Token))
}

func (u *Token) TableName() string {
	return "token"
}
