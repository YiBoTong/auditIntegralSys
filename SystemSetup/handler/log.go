package handler

import (
	"auditIntegralSys/SystemSetup/db/log"
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Log struct {
	gmvc.Controller
}

func (l *Log) List() {
	reqData := l.Request.GetJson()
	var rspData []entity.Log
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	key := search.GetString("key")
	userId := search.GetInt("userId")

	searchMap := g.Map{}

	if key != "" {
		searchMap["`key`"] = key
	}

	if userId != 0 {
		searchMap["user_id"] = userId
	}

	count, err := db_log.GetLogCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_log.GetLogs(offset, size, searchMap)
		for _, v := range listData {
			rspData = append(rspData, entity.Log{
				Id:       gconv.Int(v["id"]),
				Url:      gconv.String(v["url"]),
				Server:   gconv.String(v["server"]),
				UserId:   gconv.Int(v["user_id"]),
				UserName: gconv.String(v["user_name"]),
				Method:   gconv.String(v["method"]),
				Msg:      gconv.String(v["msg"]),
				Time:     gconv.String(v["time"]),
			})
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Log List]: %v", err)
	}
	l.Response.WriteJson(app.ListResponse{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.ListStr, err != nil),
		},
		Page: app.Pager{
			Page:  page,
			Size:  size,
			Total: count,
		},
	})
}

func (l *Log) Delete() {
	logId := l.Request.GetQueryInt("id")
	rows, err := db_log.DelLog(logId)
	if err != nil {
		log.Instance().Errorfln("[Log Delete]: %v", err)
	}
	success := err == nil && rows > 0
	l.Response.WriteJson(app.Response{
		Data: logId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
