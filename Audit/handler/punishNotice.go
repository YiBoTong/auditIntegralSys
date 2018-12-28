package handler

import (
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/db/punishNotice"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type PunishNotice struct {
	gmvc.Controller
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
		"title": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "p."+k+":"+k, gconv.String(v))
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

func (r *PunishNotice) Get() {
	id := r.Request.GetQueryInt("id")
	BasisList := []entity.PunishNoticeBasisItem{}
	Score := 0
	SumScore := 0

	PunishNoticeItem, err := db_punishNotice.Get(id)

	if err != nil {
		log.Instance().Errorfln("[PunishNotice Get]: %v", err)
	}

	if PunishNoticeItem.Id != 0 {
		basisList, _ := db_programme.GetBasis(PunishNoticeItem.DraftId)
		for _, bv := range basisList {
			item := entity.PunishNoticeBasisItem{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BasisList = append(BasisList, item)
			}
		}
	}

	success := err == nil && PunishNoticeItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.PunishNotice{
			PunishNoticeItem: PunishNoticeItem,
			Score:            Score,
			SumScore:         SumScore,
			BasisList:        BasisList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}
