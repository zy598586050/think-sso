// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Application is the golang structure for table application.
type Application struct {
	Id          int         `json:"id"          description:""`
	Name        string      `json:"name"        description:"应用名称"`
	RedirectUrl string      `json:"redirectUrl" description:"应用地址"`
	AppId       string      `json:"appId"       description:"应用ID"`
	AppSecret   string      `json:"appSecret"   description:"应用密钥"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	UpdateTime  *gtime.Time `json:"updateTime"  description:"更新时间"`
	DeleteTime  *gtime.Time `json:"deleteTime"  description:"删除时间"`
}