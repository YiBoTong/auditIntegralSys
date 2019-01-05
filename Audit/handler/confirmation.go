package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/confirmation"
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	state2 "auditIntegralSys/_public/state"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Confirmation struct {
	gmvc.Controller
}

func (r *Confirmation) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *Confirmation) checkState(state string) (bool, string) {
	hasState, msg := check.DraftState(state).Has()
	return hasState, msg
}

func (r *Confirmation) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *Confirmation) beforeState(id int, json gjson.Json) (bool, string) {
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Confirmation) stateCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	row := 0
	var err error = nil
	// 只有草稿的数据才能发布
	if stateMap["state"] == state2.Publish { // 发布
		row, err = db_confirmation.Publish(id)
	} else {
		row, err = db_confirmation.Update(id, stateMap, g.Map{"state IN (?)": g.Slice{check.D_draft}})
	}
	if err == nil && row > 0 {
		// 更新时间
		_, _ = db_confirmation.Update(id, g.Map{"update_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func (r *Confirmation) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.ConfirmationListItem
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	//search := reqData.GetJson("search")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	// todo 此处搜索需要处理
	//searchItem := map[string]interface{}{
	//	"title": "string",
	//}

	//for k, v := range searchItem {
	//	// title String
	//	util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
	//	// p.title:title String
	//	util.GetSearchMapByReqJson(listSearchMap, *search, "d."+k+":"+k, gconv.String(v))
	//}

	count, err := db_confirmation.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_confirmation.List(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.ConfirmationListItem{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Confirmation List]: %v", err)
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

func (r *Confirmation) Get() {
	id := r.Request.GetQueryInt("id")
	Draft := entity.DraftItem{}
	Programme := entity.ProgrammeItem{}
	BasisList := []entity.ProgrammeBasis{}
	DraftContent := []entity.DraftContent{}

	Confirmation, err := db_confirmation.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Confirmation Get]: %v", err)
	}

	if Confirmation.Id != 0 {
		Draft, err = db_draft.Get(Confirmation.DraftId)
		if err == nil {
			Programme, err = db_programme.Get(Draft.ProgrammeId)
		}
		if err == nil {
			basisList := g.List{}
			basisList, err = db_confirmation.GetBasis(Confirmation.Id)
			for _, cv := range basisList {
				item := entity.ProgrammeBasis{}
				if ok := gconv.Struct(cv, &item); ok == nil {
					BasisList = append(BasisList, item)
				}
			}
		}
		if err == nil {
			contentList := g.List{}
			contentList, err = db_draft.GetContent(Draft.Id)
			for _, cv := range contentList {
				item := entity.DraftContent{}
				if ok := gconv.Struct(cv, &item); ok == nil {
					DraftContent = append(DraftContent, item)
				}
			}
		}
	}

	success := err == nil && Confirmation.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Confirmation{
			ConfirmationItem: Confirmation,
			Draft:            Draft,
			Programme:        Programme,
			BasisList:        BasisList,
			DraftContent:     DraftContent,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Confirmation) Read() {
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")

	row, err := db_confirmation.Update(id,
		g.Map{
			"has_read":      1,
			"has_read_time": util.GetLocalNowTimeStr(),
		},
		// 只能更新未读的数据
		g.Map{
			"has_read": 0,
		},
	)

	if err != nil {
		log.Instance().Errorfln("[Confirmation Read]: %v", err)
	}
	success := err == nil && row != 0
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.ReadStr, !success),
		},
	})
}

// 设置依据
func (r *Confirmation) Edit() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	basisIds := reqData.GetString("basisIds")
	state := reqData.GetString("state")

	// 只有草稿的才能设置依据
	Confirmation, _ := db_confirmation.Get(id, g.Map{"state": check.D_draft})
	if Confirmation.Id != 0 {
		rows, err = db_confirmation.EditBasis(Confirmation.Id, basisIds,state)
	}
	if err != nil {
		log.Instance().Errorfln("[Draft Edit]: %v", err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}

func (r *Confirmation) State() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := r.beforeState(id, *reqData)
	if checkRes {
		rows, err = r.stateCall(id, *reqData)
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
