 package main

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/SystemSetup/handler"
	"auditIntegralSys/SystemSetup/router"
	"gitee.com/johng/gf/g"
)

const (
	apiPath = "/api/SystemSetup/"
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init()
	s := g.Server(config.SystemSetupNameSpace)
	s.BindController(apiPath+"dictionaries", new(handler.Dictionaries))
	s.BindController(apiPath+"login", new(handler.Login))
	s.BindHandler("/*", router.Index)
	s.SetLogPath(config.LogPath)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(8090)
	s.Run()
}
