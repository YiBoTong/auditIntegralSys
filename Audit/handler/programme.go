package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/programme"
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
	"time"
)

type Programme struct {
	gmvc.Controller
}

func (r *Programme) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *Programme) checkState(state string) (bool, string) {
	hasState, msg := check.ProgrammeState(state).Has()
	return hasState, msg
}

func (r *Programme) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *Programme) beforeAdd(json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canAdd, msg := r.checkState(json.GetString("state"))
	return canAdd, msg
}

func (r *Programme) addCall(json gjson.Json) (int, error) {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	reqBasis := json.GetJson("basis")
	reqContent := json.GetJson("content")
	reqStep := json.GetJson("step")
	reqBusiness := json.GetJson("business")
	reqEmphases := json.GetJson("emphases")
	reqUserList := json.GetJson("userList")

	addProgramme := map[string]interface{}{
		"title":               "string",
		"key":                 "string",
		"query_department_id": "int",
		"user_id":             "int",
		"query_point_id":      "int",
		"purpose":             "string",
		"type":                "string",
		"start_time":          "string",
		"end_time":            "string",
		"plan_start_time":     "string",
		"plan_end_time":       "string",
		"state":               "string",
		"update_time":         "nowTime", // 当前时间
	}
	addBasis := map[string]interface{}{
		"clause_id": "int",
		"order":     "int",
		"content":   "string",
	}
	addContent := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addStep := map[string]interface{}{
		"order":   "int",
		"content": "string",
		"type":    "string",
	}
	addBusiness := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addEmphases := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addUserList := map[string]interface{}{
		"user_id": "int",
		"job":     "string",
		"title":   "string",
		"task":    "string",
		"order":   "int",
	}

	programme := g.Map{
		"author_id": thisUserId,
	}
	basis := []g.Map{}
	content := []g.Map{}
	step := []g.Map{}
	business := []g.Map{}
	emphases := []g.Map{}
	userList := []g.Map{}

	util.GetSqlMap(json, addProgramme, programme)

	util.GetSqlMapItemFun(*reqBasis, addBasis, func(itemMap g.Map) {
		basis = append(basis, itemMap)
	})
	util.GetSqlMapItemFun(*reqContent, addContent, func(itemMap g.Map) {
		content = append(content, itemMap)
	})
	util.GetSqlMapItemFun(*reqStep, addStep, func(itemMap g.Map) {
		step = append(step, itemMap)
	})
	util.GetSqlMapItemFun(*reqBusiness, addBusiness, func(itemMap g.Map) {
		business = append(business, itemMap)
	})
	util.GetSqlMapItemFun(*reqEmphases, addEmphases, func(itemMap g.Map) {
		emphases = append(emphases, itemMap)
	})
	util.GetSqlMapItemFun(*reqUserList, addUserList, func(itemMap g.Map) {
		userList = append(userList, itemMap)
	})

	//fmt.Println(programme)
	//fmt.Println(basis)
	//fmt.Println(content)
	//fmt.Println(step)
	//fmt.Println(business)
	//fmt.Println(emphases)
	//fmt.Println(userList)
	id, err := db_programme.Add(programme, basis, content, step, business, emphases, userList)
	return id, err
}

func (r *Programme) beforeEdit(id int, json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Programme) editCall(id int, json gjson.Json) (int, error) {
	reqBasis := json.GetJson("basis")
	reqContent := json.GetJson("content")
	reqStep := json.GetJson("step")
	reqBusiness := json.GetJson("business")
	reqEmphases := json.GetJson("emphases")
	reqUserList := json.GetJson("userList")

	editProgramme := map[string]interface{}{
		"title":               "string",
		"key":                 "string",
		"query_department_id": "int",
		"user_id":             "int",
		"query_point_id":      "int",
		"purpose":             "string",
		"type":                "string",
		"start_time":          "string",
		"end_time":            "string",
		"plan_start_time":     "string",
		"plan_end_time":       "string",
		"state":               "string",
		"update_time":         "nowTime", // 当前时间
	}
	editBasis := map[string]interface{}{
		"id":        "int",
		"clause_id": "int",
		"order":     "int",
		"content":   "string",
	}
	editContent := map[string]interface{}{
		"id":      "int",
		"order":   "int",
		"content": "string",
	}
	editStep := map[string]interface{}{
		"id":      "int",
		"order":   "int",
		"content": "string",
		"type":    "string",
	}
	editBusiness := map[string]interface{}{
		"id":      "int",
		"order":   "int",
		"content": "string",
	}
	editEmphases := map[string]interface{}{
		"id":      "int",
		"order":   "int",
		"content": "string",
	}
	editUserList := map[string]interface{}{
		"id":      "int",
		"user_id": "int",
		"job":     "string",
		"title":   "string",
		"task":    "string",
		"order":   "int",
	}

	where := g.Map{
		// 只有草稿、部门负责人驳回两种状态可以被修改
		"state IN (?)": g.Slice{check.P_draft, check.P_dep_reject, check.P_admin_reject},
	}

	programme := g.Map{}
	basis := [2][]g.Map{}
	content := [2][]g.Map{}
	step := [2][]g.Map{}
	business := [2][]g.Map{}
	emphases := [2][]g.Map{}
	userList := [2][]g.Map{}

	rows := 0
	err := error(nil)

	util.GetSqlMap(json, editProgramme, programme)

	util.GetSqlMapItemFun(*reqBasis, editBasis, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		basis[index] = append(basis[index], itemMap)
	})
	util.GetSqlMapItemFun(*reqContent, editContent, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		content[index] = append(content[index], itemMap)
	})
	util.GetSqlMapItemFun(*reqStep, editStep, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		step[index] = append(step[index], itemMap)
	})
	util.GetSqlMapItemFun(*reqBusiness, editBusiness, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		business[index] = append(business[index], itemMap)
	})
	util.GetSqlMapItemFun(*reqEmphases, editEmphases, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		emphases[index] = append(emphases[index], itemMap)
	})
	util.GetSqlMapItemFun(*reqUserList, editUserList, func(itemMap g.Map) {
		index := 1
		if itemMap["id"] == nil {
			index = 0
		}
		userList[index] = append(userList[index], itemMap)
	})

	rows, err = db_programme.Edit(id, programme, basis, content, step, business, emphases, userList, where)

	//fmt.Println(programme)
	//fmt.Println(basis)
	//fmt.Println(content)
	//fmt.Println(step)
	//fmt.Println(business)
	//fmt.Println(emphases)
	//fmt.Println(userList)
	return rows, err
}

func (r *Programme) beforeState(id int, json gjson.Json) (bool, string) {
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Programme) stateCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	// 只有草稿的数据才能上报
	row, err := db_programme.Update(id, stateMap, g.Map{"state IN (?)": g.Slice{check.P_draft, check.P_dep_reject, check.P_admin_reject}})
	if err == nil && row > 0 {
		// 更新时间
		_, _ = db_programme.Update(id, g.Map{"update_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func (r *Programme) beforeDepExamine(id int, json gjson.Json) (bool, string) {
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Programme) depExamineCall(id int, json gjson.Json) (int, error) {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	state := map[string]interface{}{
		"state":           "string",
		"content":         "string",
		"programme_id:id": "int",
	}
	stateMap := g.Map{
		"user_id": thisUserId,
	}
	util.GetSqlMap(json, state, stateMap)
	programmeState := stateMap["state"]
	if programmeState == check.P_adopt {
		// 部门领导通过
		programmeState = check.P_dep_adopt
	} else {
		// 部门领导驳回
		programmeState = check.P_dep_reject
	}
	// 只有上报的数据才能被部门负责人进行审核
	row, err := db_programme.Update(id, g.Map{"state": programmeState}, g.Map{"state=?": check.P_report})
	if err == nil && row > 0 {
		// 更新时间
		_, _ = db_programme.AddDepExamines(id, stateMap)
	}
	return row, err
}

func (r *Programme) beforeAdminExamine(id int, json gjson.Json) (bool, string) {
	canEdit, msg := r.checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func (r *Programme) adminExamineCall(id int, json gjson.Json) (int, error) {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	state := map[string]interface{}{
		"state":           "string",
		"content":         "string",
		"programme_id:id": "int",
	}
	stateMap := g.Map{
		"user_id": thisUserId,
	}
	util.GetSqlMap(json, state, stateMap)
	programmeState := stateMap["state"]
	if programmeState == check.P_adopt {
		// 分管领导通过则直接发布
		programmeState = check.P_publish
	} else {
		// 分管领导驳回
		programmeState = check.P_admin_reject
	}
	// 只有部门负责人审核通过的分管领导才能审核
	row, err := db_programme.Update(id, g.Map{"state": programmeState}, g.Map{"state=?": check.P_dep_adopt})
	if err == nil && row > 0 {
		// 更新时间
		_, _ = db_programme.AddAdminExamines(id, stateMap)
	}
	return row, err
}

func (r *Programme) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.ProgrammeItem{}
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
		"title": "string",
		"state": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	count, err := db_programme.Count(thisUserId, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(thisUserId, offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeItem{}
			if ok := gconv.Struct(v, &programmeItem); ok == nil {
				rspData = append(rspData, programmeItem)
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

func (r *Programme) Select() {
	reqData := r.Request.GetJson()
	rspData := []entity.ProgrammeSelectItem{}
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	today := time.Now().Format("2006-01-02")

	searchMap := g.Map{
		"state":      state.Publish,
		"start_time": today,
		"end_time":   today,
	}
	listSearchMap := g.Map{
		"state":      state.Publish,
		"start_time": today,
		"end_time":   today,
	}

	searchItem := map[string]interface{}{
		"title": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	count, err := db_programme.Count(thisUserId, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(thisUserId, offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeSelectItem{}
			if ok := gconv.Struct(v, &programmeItem); ok == nil {
				rspData = append(rspData, programmeItem)
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

func (r *Programme) Add() {
	id := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	checkRes, msg := r.beforeAdd(*reqData)
	if checkRes {
		id, err = r.addCall(*reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[Programme Add]: %v", err)
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

func (r *Programme) Get() {
	id := r.Request.GetQueryInt("id")
	basis := []entity.ProgrammeBasis{}
	content := []entity.ProgrammeContent{}
	step := []entity.ProgrammeStep{}
	business := []entity.ProgrammeBusiness{}
	emphases := []entity.ProgrammeEmphases{}
	userList := []entity.ProgrammeUser{}
	depExamines := []entity.ProgrammeDepExamine{}
	adminExamines := []entity.ProgrammeAdminExamine{}

	programme, err := db_programme.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Programme Get]: %v", err)
	}

	if programme.Id != 0 {
		basisList, _ := db_programme.GetBasis(id)
		contentList, _ := db_programme.GetContent(id)
		stepList, _ := db_programme.GetStep(id)
		businessList, _ := db_programme.GetBusiness(id)
		emphasesList, _ := db_programme.GetEmphases(id)
		userListList, _ := db_programme.GetUser(id)
		depExamineList, _ := db_programme.GetDepExamines(id)
		adminExamineList, _ := db_programme.GetAdminExamines(id)

		for _, bv := range basisList {
			item := entity.ProgrammeBasis{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				basis = append(basis, item)
			}
		}

		for _, cv := range contentList {
			item := entity.ProgrammeContent{}
			if ok := gconv.Struct(cv, &item); ok == nil {
				content = append(content, item)
			}
		}

		for _, sv := range stepList {
			item := entity.ProgrammeStep{}
			if ok := gconv.Struct(sv, &item); ok == nil {
				step = append(step, item)
			}
		}

		for _, bsv := range businessList {
			item := entity.ProgrammeBusiness{}
			if ok := gconv.Struct(bsv, &item); ok == nil {
				business = append(business, item)
			}
		}

		for _, ev := range emphasesList {
			item := entity.ProgrammeEmphases{}
			if ok := gconv.Struct(ev, &item); ok == nil {
				emphases = append(emphases, item)
			}
		}

		for _, uv := range userListList {
			item := entity.ProgrammeUser{}
			if ok := gconv.Struct(uv, &item); ok == nil {
				userList = append(userList, item)
			}
		}

		for _, uv := range depExamineList {
			item := entity.ProgrammeDepExamine{}
			if ok := gconv.Struct(uv, &item); ok == nil {
				depExamines = append(depExamines, item)
			}
		}

		for _, uv := range adminExamineList {
			item := entity.ProgrammeAdminExamine{}
			if ok := gconv.Struct(uv, &item); ok == nil {
				adminExamines = append(adminExamines, item)
			}
		}
	}

	success := err == nil && programme.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Programme{
			ProgrammeItem: programme,
			Basis:         basis,
			Content:       content,
			Step:          step,
			Business:      business,
			Emphases:      emphases,
			UserList:      userList,
			DepExamines:   depExamines,
			AdminExamines: adminExamines,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Programme) Edit() {
	rows := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := r.beforeEdit(id, *reqData)
	if checkRes {
		rows, err = r.editCall(id, *reqData)
	}
	if err != nil {
		log.Instance().Errorfln("[Programme Edit]: %v", err)
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

// 创建者改变状态
func (r *Programme) State() {
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

// 部门负责人审核
func (r *Programme) Dep_examine() {
	rows := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := r.beforeDepExamine(id, *reqData)
	if checkRes {
		rows, err = r.depExamineCall(id, *reqData)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.ExamineStr, !success)
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

// 分管领导审核
func (r *Programme) Admin_examine() {
	rows := 0
	err := error(nil)
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := r.beforeAdminExamine(id, *reqData)
	if checkRes {
		rows, err = r.adminExamineCall(id, *reqData)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.ExamineStr, !success)
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

func (r *Programme) Delete() {
	id := r.Request.GetQueryInt("id")
	rows, err := db_programme.Del(id)
	if err != nil {
		log.Instance().Errorfln("[Programme Delete]: %v", err)
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
