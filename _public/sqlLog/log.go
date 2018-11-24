package SqlLog

import (
	"auditIntegralSys/SystemSetup/db/log"
	"auditIntegralSys/_public/util"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type SqlLog struct {
	key     string
	user_id int
	msg     string
	method  string
	data    string
	server  string
	ip      string
}

var logs *SqlLog

func (l *SqlLog) Type(t string) *SqlLog {
	return &SqlLog{key: t}
}

func (l *SqlLog) UserId(uId int) *SqlLog {
	return &SqlLog{user_id: uId}
}

func (l *SqlLog) Msg(msg string) *SqlLog {
	return &SqlLog{msg: msg}
}

func (l *SqlLog) Method(method string) *SqlLog {
	return &SqlLog{method: method}
}

func (l *SqlLog) Data(data interface{}) *SqlLog {
	return &SqlLog{data: gconv.String(data)}
}

func (l *SqlLog) Server(server string) *SqlLog {
	return &SqlLog{server: server}
}

func (l *SqlLog) Ip(ip string) *SqlLog {
	return &SqlLog{ip: ip}
}

func (l *SqlLog) Done() {
	_, _ = db_log.AddLog(g.Map{
		"key":     l.key,
		"user_id": l.user_id,
		"msg":     l.msg,
		"method":  l.method,
		"data":    l.data,
		"server":  l.server,
		"ip":      l.ip,
		"time":    util.GetLocalNowTimeStr(),
	})
	l.clear()
}

func (l *SqlLog) clear() *SqlLog {
	return &SqlLog{
		key:     "",
		user_id: 0,
		msg:     "",
		method:  "",
		data:    "",
		server:  "",
		ip:      "",
	}
}

func Init(controller *gmvc.Controller) *SqlLog {
	return &SqlLog{
		key:     "user",
		user_id: 1,
		method:  controller.Request.Method,
		ip:      controller.Request.GetClientIp(),
	}
}

func Instance() *SqlLog {
	return logs
}
