package model

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Id         int         `json:"id"         description:""`
	Name       string      `json:"name"       description:"用户名"`
	Phone      string      `json:"phone"      description:"手机号"`
	Email      string      `json:"email"      description:"邮箱"`
	AppIds     string      `json:"appIds"     description:"用户关联的应用"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"修改时间"`
}
