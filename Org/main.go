package main

import (
	"auditIntegralSys/Org/handler"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/router"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.OrgNameSpace
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.OrgNameSpace)
	s := g.Server(config.OrgNameSpace)
	s.BindController(apiPath+"/department", new(handler.Department))
	s.BindHandler("/*", router.Index)
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		500: router.Status_500,
	})
	s.SetLogPath(config.LogPath)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(g.Config().GetInt("appPort"))
	s.Run()
}
