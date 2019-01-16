package handler

import (
	"auditIntegralSys/Org/check"
	"auditIntegralSys/Org/db/clause"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Org/fun"
	"auditIntegralSys/Worker/db/file"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"regexp"
)

type Clause struct {
	gmvc.Controller
}

var h1Regexp = regexp.MustCompile(`^第\S{1,2}章`)
var h2Regexp = regexp.MustCompile(`^第\S{1,2}节`)

var docxRegexp = regexp.MustCompile(`^docx$`)

func (r *Clause) importCall(departmentId int, listArr g.SliceStr) (g.Map, g.List) {
	add := g.Map{
		"department_id": departmentId,
		"title":         "",
		"number":        "",
		"from":          "",
		"author_id":     util.GetUserIdByRequest(r.Cookie),
		"update_time":   util.GetLocalNowTimeStr(),
		"state":         state.Draft,
	}
	addContent := g.List{}
	for _, v := range listArr {
		if v == "" && add["from"] == "" {
			continue
		}
		if add["from"] == "" {
			add["from"] = v
			continue
		}
		if add["number"] == "" {
			add["number"] = v
			continue
		}
		if v == "" && add["title"] == "" {
			continue
		}
		if add["title"] == "" {
			add["title"] = v
			continue
		}
		titleLevel := ""
		if h1Regexp.MatchString(v) {
			titleLevel = "h1"
		}
		if titleLevel == "" && h2Regexp.MatchString(v) {
			titleLevel = "h2"
		}
		addContent = append(addContent, g.Map{
			"title_level": titleLevel,
			"is_title":    gconv.Bool(titleLevel),
			"content":     v,
		})
	}
	return add, addContent
}

func (r *Clause) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.Clause{}
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	searchMap := g.Map{}
	searchListMap := g.Map{}

	searchItem := map[string]interface{}{
		"title":         "string",
		"state":         "string",
		"c.type:type":   "type",
		"department_id": "int",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		util.GetSearchMapByReqJson(searchListMap, *search, k, gconv.String(v))
	}

	count, err := db_clause.GetClauseCount(thisUserId, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_clause.GetClauses(thisUserId, offset, size, searchListMap)
		for _, v := range listData {
			item := entity.Clause{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Notice List]: %v", err)
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

func (r *Clause) Search() {
	err := error(nil)
	content := r.Request.GetString("content")
	rspData := []entity.ClauseContent{}

	if len(content) > 1 {
		contentRes := g.List{}
		contentRes, err = db_clause.SearchClauseContents(content)
		// 查询内容
		for _, v := range contentRes {
			item := entity.ClauseContent{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}

	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.ListStr, err != nil),
		},
	})
}

func (r *Clause) Title() {
	err := error(nil)
	title := r.Request.GetString("title")
	departmentId := r.Request.GetInt("departmentId")
	rspData := []entity.ClauseTitle{}

	if len(title) > 1 {
		contentRes := g.List{}
		contentRes, err = db_clause.GetClauseTitle(0, 10, departmentId, title)
		// 查询标题
		for _, v := range contentRes {
			item := entity.ClauseTitle{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
		}
	}

	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.ListStr, err != nil),
		},
	})
}

func (r *Clause) Add() {
	reqData := r.Request.GetJson()
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	departmentId := reqData.GetInt("departmentId")
	contentList := reqData.GetJson("content")
	state := reqData.GetString("state")

	id := 0
	msg := ""
	hasState := false
	hasDepartment := false
	err := error(nil)
	if departmentId > 0 {
		// 检测是否部门是否存在
		hasDepartment, msg, err = check.HasDepartment(departmentId)
	} else {
		// 所有部门通用
		departmentId = -1
		hasDepartment = true
	}
	if !hasDepartment {
		err = errors.New(msg)
	}
	if err == nil {
		// 检测状态是否合法
		hasState, msg = check.ClauseState(state).HasState()
		if !hasState {
			err = errors.New(msg)
		}
	}
	// 添加办法
	if hasDepartment && err == nil {
		addClause := g.Map{
			"title":  "string",
			"number": "string",
			"from":   "string",
			"type":   "string",
			"state":  "string",
		}
		addClauseContent := g.Map{
			"is_title":    "int8",
			"title_level": "string",
			"content":     "string",
			"order":       "int",
		}
		clause := g.Map{
			"department_id": departmentId,
			"author_id":     thisUserId,
			"update_time":   util.GetLocalNowTimeStr(),
		}
		clauseContent := g.List{}

		util.GetSqlMap(*reqData, addClause, clause)

		util.GetSqlMapItemFun(*contentList, addClauseContent, func(itemMap g.Map) {
			clauseContent = append(clauseContent, itemMap)
		})

		id, err = db_clause.AddByTX(clause, clauseContent, reqData.GetString("fileIds"))
	}
	if err != nil && id > 0 {
		_, _ = db_clause.DelClause(id)
		log.Instance().Errorfln("[Notice Add]: %v", err)
	}
	success := err == nil && id > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.AddStr, !success)
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

func (r *Clause) Import() {
	fileId := r.Request.GetQueryInt("fileId")
	departmentId := r.Request.GetQueryInt("departmentId")
	id := 0
	msg := ""
	listArr := g.SliceStr{}
	file := entity2.File{}
	var err error
	if departmentId == 0 || departmentId == -1 {
		departmentId = -1
	} else {
		// 检测是否部门是否存在
		hasDepartment := false
		hasDepartment, msg, err = check.HasDepartment(departmentId)
		if !hasDepartment {
			err = errors.New(msg)
		}
	}
	if err == nil {
		file, err = db_file.Get(fileId)
	}
	if err == nil && file.Id != 0 {
		filePath := g.Config().GetString("filePath") + file.Path + file.FileName + "." + file.Suffix
		if docxRegexp.MatchString(file.Suffix) {
			listArr, err = fun.ReadWord(filePath)
		}
	}
	if err == nil {
		if len(listArr) > 0 {
			addData, addContents := r.importCall(departmentId, listArr)
			if len(addContents) > 0 {
				addData["type"] = "other"
				id, err = db_clause.ImportByTX(addData, addContents, fileId)
			}
		} else {
			msg = config.ImportStr + config.ErrorStr
		}
	}
	success := err == nil && id > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.ImportStr, !success)
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

func (r *Clause) Get() {
	clauseId := r.Request.GetQueryInt("id")
	fileList := []entity2.File{}
	contentList := []entity.ClauseContent{}
	clauseInfo, err := db_clause.GetClause(clauseId)

	if clauseInfo.Id > 0 && err == nil {
		var contentRes []map[string]interface{}
		// 查询内容
		contentRes, err = db_clause.GetClauseContents(clauseId, 0, 400, g.Map{})
		for _, v := range contentRes {
			item := entity.ClauseContent{}
			if ok := gconv.Struct(v, &item); ok == nil {
				contentList = append(contentList, item)
			}
		}
	}

	if clauseInfo.Id > 0 && err == nil {
		var fileRes []map[string]interface{}
		// 查询附件
		fileRes, err = db_file.GetFilesByFrom(clauseInfo.Id, table.Clause)
		for _, v := range fileRes {
			item := entity2.File{}
			if ok := gconv.Struct(v, &item); ok == nil {
				fileList = append(fileList, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Notice Get]: %v", err)
	}
	success := err == nil && clauseInfo.Id > 0
	r.Response.WriteJson(app.Response{
		Data: entity.ClauseRes{
			Clause:   clauseInfo,
			Content:  contentList,
			FileList: fileList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Clause) State() {
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	state := reqData.GetString("state")
	rows := 0
	err := error(nil)
	// 检测状态是否合法
	hasState, msg := check.ClauseState(state).HasState()
	if hasState {
		rows, err = db_clause.UpdateClause(id, g.Map{
			"state": state,
		})
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

func (r *Clause) Edit() {
	reqData := r.Request.GetJson()
	clauseId := reqData.GetInt("id")
	departmentId := reqData.GetInt("departmentId")
	contentList := reqData.GetJson("content")
	cListLen := len(contentList.ToArray())
	state := reqData.GetString("state")

	msg := ""
	rows := 0
	hasState := false
	hasDepartment := false
	err := error(nil)
	if departmentId > 0 {
		// 检测是否部门是否存在
		hasDepartment, msg, err = check.HasDepartment(departmentId)
	} else {
		// 所有部门通用
		departmentId = -1
		hasDepartment = true
	}
	if err == nil {
		// 检测状态是否合法
		hasState, msg = check.NoticeState(state).HasState()
		if !hasState {
			err = errors.New(msg)
		}
	}
	// 更新办法
	if hasDepartment && err == nil && clauseId > 0 {
		rows, err = db_clause.UpdateClause(clauseId, g.Map{
			"department_id": departmentId,
			"title":         reqData.GetString("title"),
			"from":          reqData.GetString("from"),
			"type":          reqData.GetString("type"),
			"number":        reqData.GetString("number"),
			"author_id":     util.GetUserIdByRequest(r.Cookie),
			"update_time":   util.GetLocalNowTimeStr(),
			"state":         state,
		})
	}
	// 内容处理
	if err == nil && clauseId > 0 && cListLen > 0 {
		var addContentArr []g.Map
		var updateContentArr []g.Map
		for i := 0; i < cListLen; i++ {
			id := contentList.GetInt(gconv.String(i) + ".id")
			if id != 0 {
				updateContentArr = append(updateContentArr, g.Map{
					"id":          id,
					"clause_id":   clauseId,
					"is_title":    contentList.GetBool(gconv.String(i) + ".isTitle"),
					"title_level": contentList.GetString(gconv.String(i) + ".titleLevel"),
					"content":     contentList.GetString(gconv.String(i) + ".content"),
					"order":       contentList.GetInt(gconv.String(i) + ".order"),
					"delete":      0,
				})
			} else {
				addContentArr = append(addContentArr, g.Map{
					"clause_id":   clauseId,
					"is_title":    contentList.GetBool(gconv.String(i) + ".isTitle"),
					"title_level": contentList.GetString(gconv.String(i) + ".titleLevel"),
					"content":     contentList.GetString(gconv.String(i) + ".content"),
					"order":       contentList.GetInt(gconv.String(i) + ".order"),
				})
			}
		}
		_, _ = db_clause.DelClauseContentByClauseId(clauseId)
		if len(addContentArr) > 0 {
			_, err = db_clause.AddClauseContents(addContentArr)
		}
		if err == nil && len(updateContentArr) > 0 {
			_, err = db_clause.UpdateClauseContents(updateContentArr)
		}
	}
	// 添加附件
	if err == nil && clauseId > 0 {
		//_, _ = db_clause.DelClauseFile(clauseId)
		//err = db_clause.AddClauseFiles(clauseId, reqData.GetString("fileIds"))
		_, _ = db_file.DelFilesByFrom(clauseId, table.Clause)
		_, err = db_file.UpdateFileByIds(table.Clause, reqData.GetString("fileIds"), clauseId)
	}
	if err != nil {
		log.Instance().Errorfln("[Clause Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.EditStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: clauseId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (r *Clause) Delete() {
	clauseId := r.Request.GetQueryInt("id")
	rows, err := db_clause.DelClause(clauseId)
	if err != nil {
		log.Instance().Errorfln("[Clause Delete]: %v", err)

	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: clauseId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
