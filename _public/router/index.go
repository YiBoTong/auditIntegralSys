package router

import (
	"auditIntegralSys/_public/app"
	"gitee.com/johng/gf/g/net/ghttp"
)

func Index(r *ghttp.Request) {
	r.Response.WriteJson(app.Response{
		Data: "API running",
		Status: app.Status{
			Code:  0,
			Error: false,
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