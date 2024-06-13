package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"io/ioutil"
	"net/http"
	"server/internal/service"
	"server/utility"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"错误码"`
	Message string      `json:"message" dc:"消息"`
	Data    interface{} `json:"data"    dc:"内容"`
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	excludePaths := g.Cfg().MustGet(ctx, "token.excludePaths").Strings()
	for _, p := range excludePaths {
		if gstr.Equal(r.Request.URL.Path, p) {
			r.Middleware.Next()
			return
		}
	}
	request, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8369/api/v1/check/auth", nil)
	token, err := utility.GetAuthorization(r)
	if err != nil {
		r.Response.WriteJson(&DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: err.Error(),
		})
		r.ExitAll()
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		r.Response.WriteJson(&DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: err.Error(),
		})
		r.ExitAll()
	}
	resString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		r.Response.WriteJson(&DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: err.Error(),
		})
		r.ExitAll()
	}
	resObj, err := gjson.DecodeToJson(string(resString))
	if err != nil {
		r.Response.WriteJson(&DefaultHandlerResponse{
			Code:    gcode.CodeInternalError.Code(),
			Message: err.Error(),
		})
		r.ExitAll()
	}
	if resObj.Get("code").Int() != 0 {
		r.Response.WriteJson(&DefaultHandlerResponse{
			Code:    resObj.Get("code").Int(),
			Message: resObj.Get("message").String(),
		})
		r.ExitAll()
	}
	r.Middleware.Next()
}

func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
