package controller

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"server/api/v1"
	"server/internal/model"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

func (c *cLogin) CodeLogin(ctx context.Context, req *v1.CodeLoginReq) (res *v1.LoginRes, err error) {
	response, err := g.Client().Header(g.MapStrStr{
		"Content-Type": "application/json",
	}).Post(ctx, "http://127.0.0.1:8369/api/v1/login/code", g.Map{
		"code":      req.Code,
		"appId":     "wx60adef96f04c9ed5",
		"AppSecret": "1cded5a39b1e4172406997a9b0388165",
	})
	defer response.Close()
	if err != nil {
		return
	}
	resObj, err := gjson.DecodeToJson(response.ReadAllString())
	if err != nil {
		return
	}
	if resObj.Get("code").Int() != 0 {
		err = gerror.New(resObj.Get("message").String())
		return
	} else {
		res = &v1.LoginRes{
			User: model.User{
				Id:         resObj.Get("data.id").Int(),
				Name:       resObj.Get("data.name").String(),
				Avatar:     resObj.Get("data.avatar").String(),
				Phone:      resObj.Get("data.phone").String(),
				Email:      resObj.Get("data.email").String(),
				AppIds:     resObj.Get("data.appIds").String(),
				CreateTime: resObj.Get("data.createTime").GTime(),
				UpdateTime: resObj.Get("data.updateTime").GTime(),
			},
		}
	}
	return
}

func (c *cLogin) Test(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	return
}
