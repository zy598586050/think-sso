package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id         int         `json:"id"         description:""`
	Name       string      `json:"name"       description:"用户名"`
	Avatar     string      `json:"avatar"     description:"头像"`
	Phone      string      `json:"phone"      description:"手机号"`
	Email      string      `json:"email"      description:"邮箱"`
	AppIds     string      `json:"appIds"     description:"用户关联的应用"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"修改时间"`
}

type JwtUser struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}
