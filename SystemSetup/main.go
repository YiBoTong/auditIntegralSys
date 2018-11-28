package main

import (
	"auditIntegralSys/SystemSetup/handler"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/router"
	"auditIntegralSys/_public/sqlLog"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.SystemSetupNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.SystemSetupNameSpace)
	s := g.Server(config.SystemSetupNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	s.BindController(apiPath+"/dictionaries", new(handler.Dictionaries))
	s.BindController(apiPath+"/login", new(handler.Login))
	s.BindController(apiPath+"/log", new(handler.Log))
	s.BindHandler("/*", router.Index)
	_ = s.BindHookHandlerByMap(apiPath+"/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			server := r.Server.GetName()
			userId := util.GetUserIdByRequest(r.Cookie)
			log.Instance().Debugfln("测试 %v", server)
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
	s.Run()
}
