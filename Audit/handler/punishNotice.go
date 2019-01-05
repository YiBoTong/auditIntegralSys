package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/confirmation"
	"auditIntegralSys/Audit/db/integral"
	"auditIntegralSys/Audit/db/punishNotice"
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

type PunishNotice struct {
	gmvc.Controller
}

func (r *PunishNotice) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *PunishNotice) checkState(state string) (bool, string) {
	hasState, msg := check.PublicState(state).Has()
	return hasState, msg
}

func (r *PunishNotice) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *PunishNotice) editCall(id, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	row := 0
	behaviorList := json.GetJson("behaviorList")
	BehaviorList := [2][]g.Map{}
	editBehavior := map[string]interface{}{
		"id":          "int",
		"behavior_id": "int",
		"content":     "string",
	}
	// 只有草稿状态的才能填写违规行为
	updateTime := util.GetLocalNowTimeStr()
	punishNotice, err := db_punishNotice.Get(id, g.Map{"pn.state": state.Draft})
	if punishNotice.Id != 0 {
		util.GetSqlMapItemFun(*behaviorList, editBehavior, func(itemMap g.Map) {
			index := 1
			if itemMap["id"] == nil {
				index = 0
			} else {
				itemMap["delete"] = 0
			}
			itemMap["update_time"] = updateTime
			itemMap["user_id"] = todoUserId
			itemMap["punish_notice_id"] = id
			BehaviorList[index] = append(BehaviorList[index], itemMap)
		})
		row, err = db_punishNotice.EditBehavior(id, BehaviorList)
	}
	// 如果提交状态是发布则更新状态为稽核草稿
	if punishNotice.Id != 0 && stateStr == state.Publish && err == nil {
		_, err = db_punishNotice.Update(id, g.Map{"state": "jh_" + state.Draft})
	}
	return row, err
}

func (r *PunishNotice) editScoreCall(id, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	row := 0
	score := json.GetInt("score")
	// 只有稽核草稿状态的才能编辑分数
	punishNotice, err := db_punishNotice.Get(id, g.Map{"pn.state": "jh_" + state.Draft})
	if punishNotice.Id != 0 {
		row, err = db_punishNotice.EditScore(id, todoUserId, score)
	}
	// 如果提交状态是发布则更新状态为领导待签署状态（稽核发布）
	if punishNotice.Id != 0 && stateStr == state.Publish && err == nil {
		_, err = db_punishNotice.Update(id, g.Map{"state": "ld_" + state.Draft})
	}
	return row, err
}

func (r *PunishNotice) editAuthorCall(id int, stateStr string) (int, error) {
	row := 0
	// 只有领导草稿状态的才能编辑被领导发布
	punishNotice, err := db_punishNotice.Get(id, g.Map{"pn.state": "ld_" + state.Draft})
	// 如果提交状态是发布则更新状态为办公室草稿状态
	if punishNotice.Id != 0 && stateStr == state.Publish && err == nil {
		row, err = db_punishNotice.Update(id, g.Map{"state": "bgs_" + state.Draft})
	}
	return row, err
}

func (r *PunishNotice) editNumberCall(id int, stateStr string, json gjson.Json) (int, error) {
	row := 0
	var err error = nil
	number := json.GetString("number")
	where := g.Map{"state": "bgs_" + state.Draft}
	// 如果提交状态是发布则更新状态为发布状态
	if stateStr == state.Publish {
		row, err = db_punishNotice.Publish(id, number, where)
	} else {
		// 只有草稿状态的才能填写
		row, err = db_punishNotice.Update(id, g.Map{"number": number}, where)
	}
	return row, err
}

func (r *PunishNotice) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.PunishNoticeItem
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

	count, err := db_punishNotice.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_punishNotice.List(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.PunishNoticeItem{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Programme List]: %v", err)
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

// 问责
func (r *PunishNotice) Get_accountability() {
	confirmationId := r.Request.GetQueryInt("confirmationId")
	resData := []entity.PunishNoticeAccountability{}

	list, err := db_punishNotice.GetUsersByConfirmationId(confirmationId)

	if err != nil {
		log.Instance().Errorfln("[PunishNotice Get]: %v", err)
	}

	if len(list) != 0 {
		for _, v := range list {
			item := entity.PunishNoticeAccountabilityUserItem{}
			BehaviorList := []entity.PunishNoticeAccountabilityUserBehaviorItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				behaviorList, _ := db_punishNotice.GetBehavior(item.Id)
				for _, bv := range behaviorList {
					bItem := entity.PunishNoticeAccountabilityUserBehaviorItem{}
					if ok := gconv.Struct(bv, &bItem); ok == nil {
						BehaviorList = append(BehaviorList, bItem)
					}
				}
				punishNoticeAccountability := entity.PunishNoticeAccountability{
					PunishNoticeAccountabilityUserItem: item,
					BehaviorList:                       BehaviorList,
				}
				resData = append(resData, punishNoticeAccountability)
			}
		}
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: resData,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *PunishNotice) Get() {
	id := r.Request.GetQueryInt("id")
	BasisList := []entity.PunishNoticeBasisItem{}
	BehaviorList := []entity.PunishNoticeBasisBehaviorItem{}
	Score := entity.PunishNoticeScore{}
	SumScore := 0

	PunishNoticeItem, err := db_punishNotice.Get(id)

	if err != nil {
		log.Instance().Errorfln("[PunishNotice Get]: %v", err)
	}

	if PunishNoticeItem.Id != 0 {
		basisList, _ := db_confirmation.GetBasis(PunishNoticeItem.ConfirmationId)
		for _, bv := range basisList {
			item := entity.PunishNoticeBasisItem{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BasisList = append(BasisList, item)
			}
		}
		behaviorList, _ := db_punishNotice.GetBehavior(PunishNoticeItem.Id)
		for _, bv := range behaviorList {
			item := entity.PunishNoticeBasisBehaviorItem{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BehaviorList = append(BehaviorList, item)
			}
		}

		Score, _ = db_punishNotice.GetScore(PunishNoticeItem.Id)
		SumScore, _ = db_integral.GetSumScore(PunishNoticeItem.Id, PunishNoticeItem.UserId)
	}

	success := err == nil && PunishNoticeItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.PunishNotice{
			PunishNoticeItem:  PunishNoticeItem,
			PunishNoticeScore: Score,
			SumScore:          SumScore,
			BasisList:         BasisList,
			BehaviorList:      BehaviorList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 填写违规行为
func (r *PunishNotice) Edit() {
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

// 稽核编辑分数
func (r *PunishNotice) Edit_score() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	stateStr := reqData.GetString("state")
	todoUserId := util.GetUserIdByRequest(r.Request.Cookie)
	checkRes, msg := r.checkIdAndState(id, stateStr)
	if checkRes {
		rows, err = r.editScoreCall(id, todoUserId, stateStr, *reqData)
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

// 领导发布
func (r *PunishNotice) Edit_author() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	stateStr := reqData.GetString("state")
	//todoUserId := util.GetUserIdByRequest(r.Request.Cookie)
	checkRes, msg := r.checkId(id)
	if checkRes {
		rows, err = r.editAuthorCall(id, stateStr)
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

// 办公室填写文件号后发布
func (r *PunishNotice) Edit_number() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	stateStr := reqData.GetString("state")
	checkRes, msg := r.checkId(id)
	if checkRes {
		rows, err = r.editNumberCall(id, stateStr, *reqData)
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
