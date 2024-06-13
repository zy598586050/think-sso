# THINK-SSO

该项目是一个通用的分布式单点登录系统，后端服务采用golang编写，基于Cookie+Redis的方式实现，您可以快速基于该系统打造属于您业务场景的单点登录服务。

#### 目录结构

```
.
├──demo                // 案例演示
│    ├─crossDomain     // 非同域业务演示
│    │  ├─server       // 业务系统后端服务
│    │  └─web          // 业务系统前端
│    └─sameDomain      // 同域业务演示
│        ├─server      // 业务系统后端服务
│        └─web         // 业务系统前端
├── server             // SSO后端服务
└── web                // SSO前端(统一登录系统)
```

该系统分为两种模式，同域和非同域模式，不同模式系统的流程有所不同，需要根据具体情况来改变写法。下面分别列举了两种模式的原理和具体的接入方式：

#### 非同域单点登录模式

![非同域单点登录系统](/images/1.png)

##### 1. 业务系统前端路由守卫

```
router.beforeEach((to, _, next) => {
    document.title = `${to.meta.title} - ${import.meta.env.VITE_APP_TITLE}`
    loadingBar.start()
    const token = getCookie('think-sso-token')
    const code = getQueryParam('code')
    if (token) {
        next()
    } else {
        if (code) {
            Login({ code }).then(() => {
                UserInfo().then(result => {
                    useUserStore().setUserInfo(result?.data)
                    window.location.href = window.location.origin
                })
            })
        } else {
            window.location.href = `${import.meta.env.VITE_SSO_URL}${window.location.href}`
        }
    }
})
```

##### 2. 业务系统前端统一请求

```
service.interceptors.response.use(
	(response) => {
		// 对响应数据做点什么
		const res = response?.data
		const code = response?.data?.code
		if (code === 61) {
			dialog.warning({
				title: '提示',
				content: '登录状态已过期，请重新登录',
				positiveText: '确定',
				maskClosable: false,
				closable: false,
				closeOnEsc: false,
				onPositiveClick: () => {
					useUserStore().clearUserInfo()
					location.href = `${import.meta.env.VITE_SSO_URL}${window.location.href}`
				}
			})
		} else if (code !== 0) {
			message.error(res.message)
			return Promise.reject(new Error(res.message))
		} else {
			return res
		}
	},
	(error) => {
		if (error.message.indexOf('timeout') != -1) {
			message.error('网络超时');
		} else if (error.message == 'Network Error') {
			message.error('网络连接错误');
		} else {
			if (error.response?.data) message.error(error.response.statusText);
			else message.error('接口路径找不到');
		}
		return Promise.reject(error);
	}
)
```

##### 3. 业务系统后端中间件

```
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
```

##### 4. 业务系统后端需要实现三个接口

* ``/login/code`` 用code登录
* ``/user/info`` 获取用户信息
* ``/logout`` 退出登录

这里列举一下 ``/logout`` 接口的实现，实际就是做了一下接口的转发，当然你也可以将这个些接口从前端直接调用，但是需要前端配合nginx做跨域处理才可以实现。

```
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
```