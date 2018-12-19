package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Programme struct {
	gmvc.Controller
}

func checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func checkState(state string) (bool, string) {
	hasState, msg := check.ProgrammeState(state).Has()
	return hasState, msg
}

func checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := checkId(id)
	if canEdit {
		canEdit, msg = checkState(state)
	}
	return canEdit, msg
}

func beforeAdd(json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canAdd, msg := checkState(json.GetString("state"))
	return canAdd, msg
}

func addCall(json gjson.Json) (int, error) {
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

	programme := g.Map{}
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

func beforeEdit(id int, json gjson.Json) (bool, string) {
	// 检测状态是否合法
	canEdit, msg := checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func editCall(id int, json gjson.Json) (int, error) {
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

	programme := g.Map{}
	basis := [2][]g.Map{}
	content := [2][]g.Map{}
	step := [2][]g.Map{}
	business := [2][]g.Map{}
	emphases := [2][]g.Map{}
	userList := [2][]g.Map{}

	rows := 0
	var err error = nil

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

	rows, err = db_programme.Edit(id, programme, basis, content, step, business, emphases, userList)

	//fmt.Println(programme)
	//fmt.Println(basis)
	//fmt.Println(content)
	//fmt.Println(step)
	//fmt.Println(business)
	//fmt.Println(emphases)
	//fmt.Println(userList)
	return rows, err
}

func beforeState(id int, json gjson.Json) (bool, string) {
	canEdit, msg := checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func stateCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	// 只有草稿的数据才能上报
	row, err := db_programme.Update(id, stateMap, g.Map{"state": "draft"})
	if err == nil {
		// 更新时间
		_, _ = db_programme.Update(id, g.Map{"update_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func beforeDepExamine(id int, json gjson.Json) (bool, string) {
	canEdit, msg := checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func depExamineCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state":                    "string",
		"det_user_content:content": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	// 只有上报的数据才能被部门负责人进行审核
	row, err := db_programme.Update(id, stateMap, g.Map{"state": "publish"})
	if err == nil {
		// 更新时间
		_, _ = db_programme.Update(id, g.Map{"det_user_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func beforeAdminExamine(id int, json gjson.Json) (bool, string) {
	canEdit, msg := checkIdAndState(id, json.GetString("state"))
	return canEdit, msg
}

func adminExamineCall(id int, json gjson.Json) (int, error) {
	state := map[string]interface{}{
		"state":                      "string",
		"admin_user_content:content": "string",
	}
	stateMap := g.Map{}
	util.GetSqlMap(json, state, stateMap)
	// 只有部门负责人审核通过的分管领导才能审核
	row, err := db_programme.Update(id, stateMap, g.Map{"state": "dep_adopt"})
	if err == nil {
		// 更新时间
		_, _ = db_programme.Update(id, g.Map{"admin_user_time": util.GetLocalNowTimeStr()})
	}
	return row, err
}

func (r *Programme) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.ProgrammeItem
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

	count, err := db_programme.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeItem{}
			err = gconv.Struct(v, &programmeItem)
			if err == nil {
				rspData = append(rspData, programmeItem)
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

func (r *Programme) Select() {
	reqData := r.Request.GetJson()
	var rspData []entity.ProgrammeSelectItem
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

	count, err := db_programme.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeSelectItem{}
			err = gconv.Struct(v, &programmeItem)
			if err == nil {
				rspData = append(rspData, programmeItem)
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

func (r *Programme) Add() {
	id := 0
	var err error = nil
	reqData := r.Request.GetJson()
	checkRes, msg := beforeAdd(*reqData)
	if checkRes {
		id, err = addCall(*reqData)
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
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := beforeEdit(id, *reqData)
	if checkRes {
		rows, err = editCall(id, *reqData)
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
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := beforeState(id, *reqData)
	if checkRes {
		rows, err = stateCall(id, *reqData)
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
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := beforeDepExamine(id, *reqData)
	if checkRes {
		rows, err = depExamineCall(id, *reqData)
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
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	checkRes, msg := beforeAdminExamine(id, *reqData)
	if checkRes {
		rows, err = adminExamineCall(id, *reqData)
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
