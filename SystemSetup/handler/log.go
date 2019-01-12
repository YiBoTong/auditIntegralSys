package handler

import (
	"auditIntegralSys/SystemSetup/db/log"
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Log struct {
	gmvc.Controller
}

func (l *Log) List() {
	reqData := l.Request.GetJson()
	rspData := []entity.Log{}
	thisUserId := util.GetUserIdByRequest(l.Cookie)
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"msg": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// d.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	count, err := db_log.GetLogCount(thisUserId, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_log.GetLogs(thisUserId, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.Log{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
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
