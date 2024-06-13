package utility

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

// ErrIsNil 异常抛出
func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			panic(msg[0])
		} else {
			panic(err.Error())
		}
	}
}

// GetAuthorization 从Header中获取token
func GetAuthorization(r *ghttp.Request) (token string, err error) {
	Authorization := r.Header.Get("Authorization")
	if Authorization == "" {
		return "", errors.New("请携带Authorization")
	}
	str := strings.Split(Authorization, " ")
	if len(str) != 2 || str[0] != "Bearer" {
		return "", errors.New("token格式错误")
	}
	return str[1], err
}
