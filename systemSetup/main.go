package main

import (
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/systemSetup/handler"
	"auditIntegralSys/systemSetup/router"
	"gitee.com/johng/gf/g"
)

const (
	apiPath = "/api/systemSetup/"
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init()
	s := g.Server(config.SystemSetupNameSpace)
	s.BindController(apiPath+"dictionaries", new(handler.Dictionaries))
	s.BindHandler("/*", router.Index)
	s.SetLogPath(config.LogPath)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(8090)
	s.Run()
}
