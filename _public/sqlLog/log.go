package sqlLog

import (
	"auditIntegralSys/SystemSetup/db/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

func Add(r *ghttp.Request, userId int) {
	_, _ = db_log.AddLog(g.Map{
		"url":     util.GetSqlLogURLByRequest(r),
		"user_id": userId,
		"msg":     util.GetSqlLogMsgByRequest(r),
		"method":  r.Request.Method,
		"data":    util.GetSqlLogDataByRequest(r),
		"server":  r.Server.GetName(),
		"ip":      r.GetClientIp(),
		"time":    util.GetLocalNowTimeStr(),
	})
}
