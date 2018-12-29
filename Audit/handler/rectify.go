package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/db/rectify"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Rectify struct {
	gmvc.Controller
}

func (r *Rectify) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *Rectify) checkState(state string) (bool, string) {
	hasState, msg := check.PublicState(state).Has()
	return hasState, msg
}

func (r *Rectify) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *Rectify) editCall(id, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	suggest := json.GetString("suggest")
	data := g.Map{
		"suggest":     suggest,
		"user_id":     todoUserId,
		"update_time": util.GetLocalNowTimeStr(),
	}
	// 只有草稿状态的才能填写违规行为
	row, err := db_rectify.Update(id, data, g.Map{"state": state.Draft})
	// 如果提交状态是发布则更新状态为稽核草稿
	if row != 0 && stateStr == state.Publish && err == nil {
		_, err = db_rectify.Update(id, g.Map{"state": state.Publish})
	}
	return row, err
}

func (r *Rectify) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.RectifyListItem
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
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, "d."+k+":"+k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "d."+k+":"+k, gconv.String(v))
	}

	count, err := db_rectify.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_rectify.List(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.RectifyListItem{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Rectify List]: %v", err)
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

func (r *Rectify) Get() {
	id := r.Request.GetQueryInt("id")
	DraftItem := entity.DraftItem{}
	DraftContent := []entity.DraftContent{}
	Programme := entity.ProgrammeItem{}
	ProgrammeBusiness := []entity.ProgrammeBusiness{}

	RectifyItem, err := db_rectify.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Rectify Get]: %v", err)
	}

	if RectifyItem.Id != 0 {
		DraftItem, _ = db_draft.Get(RectifyItem.DraftId)
		Programme, _ = db_programme.Get(RectifyItem.ProgrammeId)
		contentList, _ := db_draft.GetContent(RectifyItem.DraftId)
		programmeBusinessList, _ := db_programme.GetBusiness(RectifyItem.ProgrammeId)
		for _, bv := range contentList {
			item := entity.DraftContent{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				DraftContent = append(DraftContent, item)
			}
		}
		for _, bv := range programmeBusinessList {
			item := entity.ProgrammeBusiness{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				ProgrammeBusiness = append(ProgrammeBusiness, item)
			}
		}
	}

	success := err == nil && RectifyItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Rectify{
			RectifyItem:       RectifyItem,
			Programme:         Programme,
			Draft:             DraftItem,
			DraftContent:      DraftContent,
			ProgrammeBusiness: ProgrammeBusiness,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Rectify) Edit() {
	rows := 0
	var err error = nil
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
