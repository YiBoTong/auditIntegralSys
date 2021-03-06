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
	"strings"
)

const (
	apiPath = "/api/" + config.WorkerNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.WorkerNameSpace)
	g.DB().SetDebug(true)
	s := g.Server(config.WorkerNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	_ = s.BindController(apiPath+"/user", new(handler.User))
	_ = s.BindController(apiPath+"/file", new(handler.File))
	_ = s.BindController(apiPath+"/menu", new(handler.Menu))
	_ = s.BindHandler("/*", router.Index)
	_ = s.BindHookHandlerByMap(apiPath+"/*", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			server := r.Server.GetName()
			userId := util.GetUserIdByRequest(r.Cookie)
			log.Instance().Debugfln("server %v, userId %v, method %v, url %v", server, userId, r.Method, r.URL)
			if userId == 0 {
				if strings.Split(r.RequestURI, "?")[0] != "/api/worker/user/login" {
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
