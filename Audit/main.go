package main

import (
	"auditIntegralSys/Audit/handler"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/router"
	"auditIntegralSys/_public/sqlLog"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.AuditNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.AuditNameSpace)
	g.DB().SetDebug(true)
	s := g.Server(config.AuditNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	// 审查方案
	_ = s.BindController(apiPath+"/programme", new(handler.Programme))
	// 工作底稿
	_ = s.BindController(apiPath+"/draft", new(handler.Draft))
	// 事实确认书
	_ = s.BindController(apiPath+"/confirmation", new(handler.Confirmation))
	// 惩罚通知书
	_ = s.BindController(apiPath+"/punishNotice", new(handler.PunishNotice))
	// 整改通知
	_ = s.BindController(apiPath+"/rectify", new(handler.Rectify))
	// 整改报告
	_ = s.BindController(apiPath+"/rectifyReport", new(handler.RectifyReport))
	// 积分
	_ = s.BindController(apiPath+"/integral", new(handler.Integral))
	// 统计分析
	_ = s.BindController(apiPath+"/statistical", new(handler.Statistical))
	_ = s.BindHandler("/*", router.Index)
	_ = s.BindHookHandlerByMap(apiPath+"/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			server := r.Server.GetName()
			userId := util.GetUserIdByRequest(r.Cookie)
			log.Instance().Debugfln("server %v, userId %v, method %v, url %v", server, userId, r.Method, r.URL)
			if userId == 0 {
				router.LoginTips(r)
			} else {
				sqlLog.Add(r, userId)
			}
		},
	})
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		500: router.Status_500,
	})
	s.SetLogPath(config.LogPath)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(g.Config().GetInt("appPort"))
	_ = s.Run()
}
