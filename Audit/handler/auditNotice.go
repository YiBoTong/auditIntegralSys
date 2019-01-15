package handler

import (
	"auditIntegralSys/Audit/db/auditNotice"
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type AuditNotice struct {
	gmvc.Controller
}

func (r *AuditNotice) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.AuditNoticeListItem{}
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
		"project_name":        "string",
		"query_department_id": "int",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	thisUserInfo, _ := db_user.GetUser(thisUserId)
	count, err := db_auditNotice.Count(thisUserInfo, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_auditNotice.List(thisUserInfo, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.AuditNoticeListItem{}
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

func (r *AuditNotice) Get() {
	id := r.Request.GetQueryInt("id")
	Draft := entity.DraftItem{}
	UserList := []entity.DraftQueryUser{}
	Business := []entity.ProgrammeBusiness{}

	AuditNoticeItem, err := db_auditNotice.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Confirmation Get]: %v", err)
	}

	if AuditNoticeItem.Id != 0 {
		Draft, err = db_draft.Get(AuditNoticeItem.DraftId)
		if err == nil {
			userList := g.List{}
			userList, _ = db_draft.GetQueryUser(AuditNoticeItem.DraftId)
			for _, cv := range userList {
				item := entity.DraftQueryUser{}
				if ok := gconv.Struct(cv, &item); ok == nil {
					UserList = append(UserList, item)
				}
			}
			business, _ := db_programme.GetBusiness(Draft.ProgrammeId)
			for _, cv := range business {
				item := entity.ProgrammeBusiness{}
				if ok := gconv.Struct(cv, &item); ok == nil {
					Business = append(Business, item)
				}
			}
		}
	}

	success := err == nil && AuditNoticeItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.AuditNotice{
			AuditNoticeItem: AuditNoticeItem,
			Draft:           Draft,
			UserList:        UserList,
			Business:        Business,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}
