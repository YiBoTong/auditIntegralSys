package router

import (
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g/net/ghttp"
)

func Index(r *ghttp.Request) {
	r.Response.WriteJson(app.Response{
		Data: "API running",
		Status: app.Status{
			Code:  0,
			Error: true,
			Msg:   "URL参数不正确",
		},
	})
}

func Status_500(r *ghttp.Request)  {
	r.Response.WriteJson(app.Response{
		Data: nil,
		Status: app.Status{
			Code:  0,
			Error: true,
			Msg:   "服务异常，请重试",
		},
	})
}

func LoginTips(r *ghttp.Request)  {
	r.Cookie.Remove(config.CookieIdName,"","/")
	r.Response.WriteJson(app.Response{
		Data: "",
		Status: app.Status{
			Code:  1,
			Error: true,
			Msg:   config.LoginTispStr,
		},
	})
	r.Exit()
}