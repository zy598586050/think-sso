package model

type Application struct {
	Id   int    `json:"id"         description:""`
	Name string `json:"name"       description:"应用名称"`
	Url  string `json:"url"        description:"应用地址"`
}
