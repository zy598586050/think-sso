package controller

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"think-sso/api/v1"
	"think-sso/internal/service"
	"think-sso/utility"
	"time"
)

var (
	Login = cLogin{}
)

type cLogin struct{}

// EmailLogin 邮箱登录
func (c *cLogin) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.LoginRes, err error) {
	user, err := service.User().GetUserByEmailPassword(ctx, req)
	if err != nil {
		return
	}
	err = service.Token().CreateToken(ctx, user)
	return
}

// Code 生成Code
func (c *cLogin) Code(ctx context.Context, req *v1.CodeReq) (res *v1.CodeRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
	salt := g.Cfg().MustGet(ctx, "token.salt").String()
	// 加盐混淆加密
	code, _ := gmd5.EncryptString(token + salt)
	ttl, err := g.Redis().TTL(ctx, prefix+token)
	if err != nil {
		return
	}
	g.Redis().SetEX(ctx, prefix+code, token, ttl)
	res = &v1.CodeRes{
		Code: code,
	}
	return
}

// CodeLogin 非同域的时候code登录
func (c *cLogin) CodeLogin(ctx context.Context, req *v1.CodeLoginReq) (res *v1.LoginRes, err error) {
	_, err = service.Application().HasApp(ctx, req.AppId, req.AppSecret)
	if err != nil {
		return
	}
	prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
	token, err := g.Redis().Get(ctx, prefix+req.Code)
	if err != nil {
		return
	}
	if token.String() == "" {
		err = gerror.New("code无效")
		return
	}
	ttl, _ := g.Redis().TTL(ctx, prefix+req.Code)
	g.Redis().Del(ctx, prefix+req.Code)
	g.RequestFromCtx(ctx).Cookie.SetCookie("think-sso-token", token.String(), "", "/", time.Duration(ttl)*time.Second)
	return
}

// Logout 退出登录
func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	token, _ := utility.GetAuthorization(g.RequestFromCtx(ctx))
	err = service.Token().RemoveToken(ctx, token)
	return
}

// CheckAuth 其他系统接入，验证登录状态
func (c *cLogin) CheckAuth(ctx context.Context, req *v1.CheckAuthReq) (res *v1.CheckAuthRes, err error) {
	return
}
