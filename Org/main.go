package main

import (
	"auditIntegralSys/Org/handler"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/router"
	"auditIntegralSys/_public/sqlLog"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.OrgNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.OrgNameSpace)
	g.DB().SetDebug(true)
	s := g.Server(config.OrgNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	_ = s.BindController(apiPath+"/department", new(handler.Department))
	_ = s.BindController(apiPath+"/user", new(handler.User))
	_ = s.BindController(apiPath+"/notice", new(handler.Notice))
	_ = s.BindController(apiPath+"/clause", new(handler.Clause))
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
	s.Run()
}
