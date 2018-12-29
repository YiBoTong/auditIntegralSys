package handler

import (
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/db/rectify"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Rectify struct {
	gmvc.Controller
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
