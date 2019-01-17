package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/auditReport"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type AuditReport struct {
	gmvc.Controller
}

func (r *AuditReport) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *AuditReport) checkState(state string) (bool, string) {
	hasState, msg := check.PublicState(state).Has()
	return hasState, msg
}

func (r *AuditReport) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *AuditReport) editCall(id, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	basicInfo := json.GetString("basicInfo")
	reason := json.GetString("reason")
	plan := json.GetString("plan")
	data := g.Map{
		"state":       stateStr,
		"update_time": util.GetLocalNowTimeStr(),
	}
	// 只能更新草稿状态的数据
	row, err := db_auditReport.Edit(id, basicInfo, reason, plan, data)
	return row, err
}

func (r *AuditReport) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.AuditReportListItem{}
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	thisUserId := util.GetUserIdByRequest(r.Cookie)

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"project_name": "string",
		"state":        "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	thisUserInfo, _ := db_user.GetUser(thisUserId)
	count, err := db_auditReport.Count(thisUserInfo, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_auditReport.List(thisUserInfo, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.AuditReportListItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Draft List]: %v", err)
	}
	r.Response.WriteJson(app.ListResponse{
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

func (r *AuditReport) Get() {
	id := r.Request.GetQueryInt("id")
	BasicInfo := entity.AuditReportContent{}
	Reason := entity.AuditReportContent{}
	Plan := entity.AuditReportContent{}

	AuditReportItem, err := db_auditReport.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Draft Get]: %v", err)
	}

	if AuditReportItem.Id != 0 {
		BasicInfo, err = db_auditReport.GetBasicInfo(AuditReportItem.Id)
		Reason, err = db_auditReport.GetReason(AuditReportItem.Id)
		Plan, err = db_auditReport.GetPlan(AuditReportItem.Id)
	}

	success := err == nil && AuditReportItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.AuditReport{
			AuditReportItem: AuditReportItem,
			BasicInfo:       BasicInfo.Content,
			Reason:          Reason.Content,
			Plan:            Plan.Content,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *AuditReport) Edit() {
	rows := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	stateStr := reqData.GetString("state")
	todoUserId := util.GetUserIdByRequest(r.Request.Cookie)
	checkRes, msg := r.checkIdAndState(id, stateStr)
	if checkRes {
		rows, err = r.editCall(id, todoUserId, stateStr, *reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[PunishNotice Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.EditStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}
