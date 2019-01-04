package handler

import (
	"auditIntegralSys/Audit/db/AuditReport"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/db/statistical"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Statistical struct {
	gmvc.Controller
}

func (r *Statistical) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.StatisticalListItem
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
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "d."+k+":"+k, gconv.String(v))
	}

	count, err := db_auditReport.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_auditReport.List(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.StatisticalListItem{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
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

func (r *Statistical) Get() {
	id := r.Request.GetQueryInt("id")
	BusinessList := []entity.ProgrammeBusiness{}

	StatisticalListItem, err := db_statistical.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Rectify Get]: %v", err)
	}

	if StatisticalListItem.Id != 0 {
		basisList, _ := db_programme.GetBusiness(StatisticalListItem.ProgrammeId)
		for _, bv := range basisList {
			item := entity.ProgrammeBusiness{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BusinessList = append(BusinessList, item)
			}
		}
	}

	success := err == nil && StatisticalListItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Statistical{
			StatisticalListItem: StatisticalListItem,
			BusinessList:        BusinessList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}
