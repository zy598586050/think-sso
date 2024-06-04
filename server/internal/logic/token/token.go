package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	v1 "think-sso/api/v1"
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

func (s *sToken) CreateToken(ctx context.Context, user *v1.LoginRes) (encryptToken string, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.Id,
			"name":  user.Name,
			"phone": user.Phone,
			"email": user.Email,
			"exp":   jwt.NewNumericDate(time.Now().Add(g.Cfg().MustGet(ctx, "jwt.exp").Duration() * time.Minute)),
			"iat":   jwt.NewNumericDate(time.Now()),
		})
		token, err := t.SignedString([]byte(g.Cfg().MustGet(ctx, "jwt.signe").String()))
		if err != nil {
			utility.ErrIsNil(ctx, gerror.New("生成Token失败"))
		}
		encryptToken = gbase64.EncodeToString([]byte(gmd5.MustEncrypt(token)))
		g.Redis().SetEX(ctx, encryptToken, token, g.Cfg().MustGet(ctx, "jwt.exp").Int64()*60*60)
	})
	return
}

func (s *sToken) ValidateToken(ctx context.Context, token string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		rToken, err := g.Redis().Get(ctx, token)
		if err != nil {
			utility.ErrIsNil(ctx, gerror.New("Redis获取token失败"))
		}
		fmt.Println("看一张", rToken)
	})
	return
}
