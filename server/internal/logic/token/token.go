package token

import (
	"context"
	"crypto/ecdsa"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"think-sso/internal/model"
	"think-sso/internal/service"
	"think-sso/utility"
	"time"
)

var (
	key *ecdsa.PrivateKey
	t   *jwt.Token
)

type sToken struct{}

func init() {
	service.RegisterToken(New())
}

func New() *sToken {
	return &sToken{}
}

// CreateToken 创建Token并写入Cookie
func (s *sToken) CreateToken(ctx context.Context, user *model.User) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":      user.Id,
			"warning": "看什么看,没事儿干别乱解析",
			"exp":     jwt.NewNumericDate(time.Now().Add(g.Cfg().MustGet(ctx, "token.exp").Duration() * time.Minute)),
			"iat":     jwt.NewNumericDate(time.Now()),
		})
		token, err := t.SignedString([]byte(g.Cfg().MustGet(ctx, "token.signe").String()))
		utility.ErrIsNil(ctx, err, "生成Token失败")
		userJson, err := gjson.EncodeString(&user)
		utility.ErrIsNil(ctx, err, "user结构体转换失败")
		prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
		g.Redis().SetEX(ctx, prefix+token, userJson, g.Cfg().MustGet(ctx, "token.exp").Int64()*60)
		domain := g.Cfg().MustGet(ctx, "cookie.domain").String()
		g.RequestFromCtx(ctx).Cookie.SetCookie("think-sso-token", token, domain, "/", time.Duration(g.Cfg().MustGet(ctx, "token.exp").Int())*time.Minute)
	})
	return
}

// ParseJwt 验证并解析JWT
func (s *sToken) ParseJwt(ctx context.Context, token string) (*model.JwtUser, error) {
	j, err := jwt.ParseWithClaims(token, &model.JwtUser{}, func(tk *jwt.Token) (interface{}, error) {
		return []byte(g.Cfg().MustGet(ctx, "token.signe").String()), nil
	})
	if claims, ok := j.Claims.(*model.JwtUser); ok && j.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// CheckToken 验证登录状态
func (s *sToken) CheckToken(ctx context.Context, token string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
		userJson, err := g.Redis().Get(ctx, prefix+token)
		utility.ErrIsNil(ctx, err, "Redis获取token失败")
		if userJson.String() == "" {
			utility.ErrIsNil(ctx, gerror.New("token失效请重新登录"))
		} else {
			// 过期前10分钟内自动刷新Token
			ttl, err := g.Redis().TTL(ctx, prefix+token)
			utility.ErrIsNil(ctx, err, "Redis获取过期时间失败")
			refresh := g.Cfg().MustGet(ctx, "token.refresh").Int64()
			if ttl <= refresh*60 && g.RequestFromCtx(ctx).Request.URL.Path != "/api/v1/check/auth" {
				user, err := gjson.DecodeToJson(userJson.String())
				utility.ErrIsNil(ctx, err, "Redis解析用户失败")
				u, err := service.User().GetUserById(ctx, user.Get("id").Int())
				utility.ErrIsNil(ctx, err, "查询用户失败")
				apps, _ := service.Application().AppListByIds(ctx, u.AppIds)
				u.Apps = apps
				s.CreateToken(ctx, u)
				s.RemoveToken(ctx, token)
			}
		}
	})
	return
}

// RemoveToken 移除Token
func (s *sToken) RemoveToken(ctx context.Context, token string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
		_, err = g.Redis().Del(ctx, prefix+token)
		utility.ErrIsNil(ctx, err, "退出登录失败")
	})
	return
}

// TokenToUser 通过Token获取用户信息
func (s *sToken) TokenToUser(ctx context.Context) (res *model.User, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		token, err := utility.GetAuthorization(g.RequestFromCtx(ctx))
		utility.ErrIsNil(ctx, err, "解析用户信息失败")
		prefix := g.Cfg().MustGet(ctx, "token.prefix").String()
		userJson, err := g.Redis().Get(ctx, prefix+token)
		utility.ErrIsNil(ctx, err, "解析用户信息失败")
		user, err := gjson.DecodeToJson(userJson.String())
		utility.ErrIsNil(ctx, err, "解析用户信息失败")
		res, err = service.User().GetUserById(ctx, user.Get("id").Int())
		utility.ErrIsNil(ctx, err, "解析用户信息失败")
	})
	return
}
