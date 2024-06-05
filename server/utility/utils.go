package utility

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"math/rand"
	"strings"
	"time"
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

// GenerateSMSCode 生成短信验证码
func GenerateSMSCode(length int) int {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器
	code := 0
	multiplier := 1
	for i := 0; i < length; i++ {
		digit := rand.Intn(10) // 生成0到9的数字
		code += digit * multiplier
		multiplier *= 10
	}
	return code
}
