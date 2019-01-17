package handler

import (
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/introduction"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"time"
)

type Introduction struct {
	gmvc.Controller
}

func (r *Introduction) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // ID必须有
		canEdit = false
	}
	return canEdit, msg
}

func (r *Introduction) addCall(draftId int) (int, error) {
	id := 0
	err := error(nil)
	draft, _ := db_draft.Get(draftId)
	if draft.Id != 0 && draft.State == state.Publish {
		pre, _ := db_introduction.GetLast()
		year := time.Now().Year()
		number := fun.CreateNumber(pre.Year, pre.Number)
		introduction := g.Map{
			"draft_id": draft.Id,
			"year":     year,
			"number":   number,
		}
		id, err = db_introduction.Add(introduction)
	}
	return id, err
}

func (r *Introduction) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.IntroductionListItem{}
	thisUserId := util.GetUserIdByRequest(r.Cookie)
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
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	count, err := db_introduction.GetCount(thisUserId, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_introduction.GetList(thisUserId, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.IntroductionListItem{}
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

func (r *Introduction) Create() {
	id := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	draftId := reqData.GetInt("draftId")
	checkRes, msg := r.checkId(draftId)
	if checkRes {
		id, err = r.addCall(draftId)
	}
	if err != nil {
		log.Instance().Errorfln("[Introduction Add]: %v", err)
	}
	success := err == nil && id != 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.CreateStr, !success)
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

func (r *Introduction) Get() {
	id := r.Request.GetQueryInt("id")
	UserList := []entity.DraftQueryUser{}
	Draft := entity.DraftItem{}

	IntroductionItem, err := db_introduction.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Draft Get]: %v", err)
	}

	if IntroductionItem.Id != 0 && err == nil {
		Draft, err = db_draft.Get(IntroductionItem.DraftId)
	}
	if err == nil && Draft.Id != 0 {
		userList := g.List{}
		userList, err = db_draft.GetQueryUser(Draft.Id)
		for _, fv := range userList {
			item := entity.DraftQueryUser{}
			if ok := gconv.Struct(fv, &item); ok == nil {
				UserList = append(UserList, item)
			}
		}
	}

	success := err == nil && IntroductionItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Introduction{
			IntroductionItem: IntroductionItem,
			UserList:         UserList,
			Draft:            Draft,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}
