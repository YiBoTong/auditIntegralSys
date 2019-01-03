package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/integral"
	"auditIntegralSys/Audit/db/punishNotice"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/util"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Integral struct {
	gmvc.Controller
}

func (r *Integral) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *Integral) checkState(state string) (bool, string) {
	hasState, msg := check.PublicState(state).Has()
	return hasState, msg
}

func (r *Integral) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *Integral) beforeEdit(id int, json gjson.Json) (bool, string) {
	// 检测状态是否合法
	stateStr := json.GetString("state")
	canEdit, msg := r.checkId(id)
	if canEdit && !(stateStr != state.Draft || stateStr != state.Report) {
		msg = config.StateStr + config.NoHad
	}
	return canEdit, msg
}

func (r *Integral) beforeEditAuthor(id int, json gjson.Json) (bool, string) {
	// 检测状态是否合法
	stateStr := json.GetString("state")
	canEdit, msg := r.checkId(id)
	if canEdit && !(stateStr != state.Reject || stateStr != state.Adopt) {
		msg = config.StateStr + config.NoHad
	}
	return canEdit, msg
}

func (r *Integral) editCall(id int, json gjson.Json) (int, error) {
	add := map[string]interface{}{
		"score":       "int",
		"describe":    "string",
		"state":       "string",
		"update_time": "nowTime", // 当前时间
	}

	data := g.Map{}
	util.GetSqlMap(json, add, data)

	fmt.Println(data)

	rows, err := db_integral.ChangeScore(id, data)

	return rows, err
}

func (r *Integral) editAuthorCall(changeScoreId int, json gjson.Json) (int, error) {
	state := json.GetString("state")
	suggestion := json.GetString("suggestion")
	rows, err := db_integral.AdoptChangeScore(changeScoreId, state, suggestion)
	return rows, err
}

func (r *Integral) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.IntegralListItem
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"project_name": "string",
		"state":        "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, "i."+k+":"+k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "i."+k+":"+k, gconv.String(v))
	}

	count, err := db_integral.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_integral.List(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.IntegralListItem{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Integral List]: %v", err)
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

func (r *Integral) Get() {
	id := r.Request.GetQueryInt("id")
	IntegralItem := entity.IntegralItem{}
	SumScore := 0
	ChangeScore := entity.IntegralChangeScore{}
	BehaviorList := []entity.PunishNoticeBasisBehaviorItem{}

	IntegralItem, err := db_integral.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Rectify Get]: %v", err)
	}

	if IntegralItem.Id != 0 {
		ChangeScore, err = db_integral.GetChangeScore(IntegralItem.Id)
		SumScore, err = db_integral.GetSumScore(IntegralItem.PunishNoticeId, IntegralItem.UserId)
		behaviorList, _ := db_punishNotice.GetBehavior(IntegralItem.PunishNoticeId)
		for _, bv := range behaviorList {
			item := entity.PunishNoticeBasisBehaviorItem{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BehaviorList = append(BehaviorList, item)
			}
		}
	}

	success := err == nil && IntegralItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Integral{
			IntegralItem: IntegralItem,
			ChangeScore:  ChangeScore,
			SumScore:     SumScore,
			BehaviorList: BehaviorList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Integral) Edit() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := r.beforeEdit(id, *reqData)
	if checkRes {
		rows, err = r.editCall(id, *reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[Draft Edit]: %v", err)
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

func (r *Integral) Edit_author() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	changeScoreId := reqData.GetInt("changeScoreId")
	checkRes, msg := r.beforeEditAuthor(changeScoreId, *reqData)
	if checkRes {
		rows, err = r.editAuthorCall(changeScoreId, *reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[Draft Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.EditStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: changeScoreId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}
