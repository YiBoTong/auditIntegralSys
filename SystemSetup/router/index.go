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
