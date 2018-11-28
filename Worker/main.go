package main

import (
	"auditIntegralSys/Worker/handler"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/router"
	"auditIntegralSys/_public/sqlLog"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.WorkerNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.WorkerNameSpace)
	s := g.Server(config.WorkerNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	_ = s.BindController(apiPath+"/user", new(handler.User))
	_ = s.BindHandler("/*", router.Index)
	_ = s.BindHookHandlerByMap(apiPath+"/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			server := r.Server.GetName()
			userId := util.GetUserIdByRequest(r.Cookie)
			log.Instance().Debugfln("测试 %v", server)
			if userId == 0 {
				if r.RequestURI != "/api/worker/user/login" {
					router.LoginTips(r)
				}
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
