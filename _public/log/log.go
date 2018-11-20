package log

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g/os/glog"
)

var log *glog.Logger

func Init() {
	log = glog.New()
	log.SetPath(config.LogPath)
}

func Instance() *glog.Logger {
	return log
}
