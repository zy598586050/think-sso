package controller

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"server/api/v1"
	"server/utility"
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
		"AppSecret": "1cded5a39b1e4179406997a9b0388165",
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
		g.RequestFromCtx(ctx).Response.Header().Set("Set-Cookie", response.Header.Get("Set-Cookie"))
	}
	return
}

func (c *cLogin) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	response, err := g.Client().Header(g.MapStrStr{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}).Post(ctx, "http://127.0.0.1:8369/api/v1/user/info", nil)
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
		err = resObj.Get("data").Scan(&res)
	}
	return
}

func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	response, err := g.Client().Header(g.MapStrStr{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}).Post(ctx, "http://127.0.0.1:8369/api/v1/logout", nil)
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
	}
	return
}
