package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/Worker/db/file"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Draft struct {
	gmvc.Controller
}

func (r *Draft) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *Draft) checkState(state string) (bool, string) {
	hasState, msg := check.DraftState(state).Has()
	return hasState, msg
}

func (r *Draft) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *Draft) beforeAdd(json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canAdd, msg := r.checkState(json.GetString("state"))
	return canAdd, msg
}

func (r *Draft) addCall(json gjson.Json) (int, error) {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	reqContent := json.GetJson("contentList")
	queryUsers := json.GetString("queryUsers")
	queryUserLeader := json.GetInt("queryUserLeader")
	adminUsers := json.GetString("adminUsers")
	fileIds := json.GetString("fileIds")

	add := map[string]interface{}{
		"project_name":        "string",
		"programme_id":        "int",
		"query_department_id": "int",
		"department_id":       "int",
		"number":              "string",
		"public":              "uint8",
		"query_start_time":    "string",
		"query_end_time":      "string",
		"state":               "string",
		"update_time":         "nowTime", // 当前时间
	}

	addContentList := map[string]interface{}{
		"order":            "int",
		"type":             "string",
		"behavior_id":      "int",
		"behavior_content": "string",
	}

	draft := g.Map{
		"author_id": thisUserId,
	}
	contentList := []g.Map{}

	util.GetSqlMap(json, add, draft)

	util.GetSqlMapItemFun(*reqContent, addContentList, func(itemMap g.Map) {
		contentList = append(contentList, itemMap)
	})

	fmt.Println(draft)
	fmt.Println(contentList)
	fmt.Println(adminUsers)
	fmt.Println(queryUsers)

	id, err := db_draft.Add(draft, contentList, queryUserLeader, queryUsers, adminUsers, fileIds)
	return id, err
}

func (r *Draft) beforeState(id int, json gjson.Json) (bool, string) {
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Draft) stateCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	row := 0
	err := error(nil)
	// 只有草稿的数据才能上报
	if stateMap["state"] == check.D_publish { // 发布
		row, err = db_draft.Publish(id)
	} else {
		row, err = db_draft.Update(id, stateMap, g.Map{"state IN (?)": g.Slice{check.D_draft}})
	}
	if err == nil && row > 0 {
		// 更新时间
		_, _ = db_draft.Update(id, g.Map{"update_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func (r *Draft) beforeEdit(id int, json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Draft) editCall(id int, json gjson.Json) (int, error) {
	reqContent := json.GetJson("contentList")
	queryUsers := json.GetString("queryUsers")
	queryUserLeader := json.GetInt("queryUserLeader")
	adminUsers := json.GetString("adminUsers")
	fileIds := json.GetString("fileIds")

	add := map[string]interface{}{
		"project_name":        "string",
		"programme_id":        "int",
		"query_department_id": "int",
		"department_id":       "int",
		"number":              "string",
		"public":              "uint8",
		"query_start_time":    "string",
		"query_end_time":      "string",
		"state":               "string",
		"update_time":         "nowTime", // 当前时间
	}

	addContentList := map[string]interface{}{
		"id":               "int",
		"order":            "int",
		"type":             "string",
		"behavior_id":      "int",
		"behavior_content": "string",
	}

	where := g.Map{
		// 只有草稿状态可以被修改
		"state IN (?)": g.Slice{state.Draft},
	}

	draft := g.Map{}
	contentList := [2][]g.Map{}

	util.GetSqlMap(json, add, draft)

	util.GetSqlMapItemFun(*reqContent, addContentList, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		contentList[index] = append(contentList[index], itemMap)
	})

	fmt.Println(draft)
	fmt.Println(contentList)
	fmt.Println(adminUsers)
	fmt.Println(queryUsers)

	rows, err := db_draft.Edit(id, draft, contentList, queryUserLeader, queryUsers, adminUsers, fileIds, where)

	return rows, err
}

func (r *Draft) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.DraftItem{}
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
		"state":               "string",
		"query_department_id": "int",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	thisUserInfo, _ := db_user.GetUser(thisUserId)
	count, err := db_draft.Count(thisUserInfo, searchMap)
	if err == nil && offset <= count {
		listData := g.List{}
		listData, err = db_draft.List(thisUserInfo, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.DraftItem{}
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

func (r *Draft) Add() {
	id := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	checkRes, msg := r.beforeAdd(*reqData)
	if checkRes {
		id, err = r.addCall(*reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[Draft Add]: %v", err)
	}
	if msg == "" {
		msg = config.GetTodoResMsg(config.AddStr, err != nil)
	}
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   msg,
		},
	})
}

func (r *Draft) Get() {
	id := r.Request.GetQueryInt("id")
	ContentList := []entity.DraftContent{}
	AdminUserList := []entity.DraftAdminUser{}
	QueryUserList := []entity.DraftQueryUser{}
	ReviewUserList := []entity.DraftReviewUser{}
	FileList := []entity2.File{}

	draft, err := db_draft.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Draft Get]: %v", err)
	}

	if draft.Id != 0 {
		contentList, _ := db_draft.GetContent(id)
		adminUserList, _ := db_draft.GetAdminUser(id)
		queryUserList, _ := db_draft.GetQueryUser(id)
		reviewUserList, _ := db_draft.GetReviewUser(id)
		fileList, _ := db_file.GetFilesByFrom(id, table.Draft)

		for _, cv := range contentList {
			item := entity.DraftContent{}
			if ok := gconv.Struct(cv, &item); ok == nil {
				ContentList = append(ContentList, item)
			}
		}

		for _, av := range adminUserList {
			item := entity.DraftAdminUser{}
			if ok := gconv.Struct(av, &item); ok == nil {
				AdminUserList = append(AdminUserList, item)
			}
		}

		for _, qv := range queryUserList {
			item := entity.DraftQueryUser{}
			if ok := gconv.Struct(qv, &item); ok == nil {
				QueryUserList = append(QueryUserList, item)
			}
		}

		for _, rv := range reviewUserList {
			item := entity.DraftReviewUser{}
			if ok := gconv.Struct(rv, &item); ok == nil {
				ReviewUserList = append(ReviewUserList, item)
			}
		}

		for _, fv := range fileList {
			item := entity2.File{}
			if ok := gconv.Struct(fv, &item); ok == nil {
				FileList = append(FileList, item)
			}
		}
	}

	success := err == nil && draft.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Draft{
			DraftItem:      draft,
			ContentList:    ContentList,
			AdminUserList:  AdminUserList,
			QueryUserList:  QueryUserList,
			ReviewUserList: ReviewUserList,
			FileList:       FileList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 创建者改变状态
func (r *Draft) State() {
	rows := 0
	err := error(nil)
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

func (r *Draft) Edit() {
	rows := 0
	err := error(nil)
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

func (r *Draft) Delete() {
	id := r.Request.GetQueryInt("id")
	rows, err := db_draft.Del(id)
	if err != nil {
		log.Instance().Errorfln("[Draft Delete]: %v", err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
