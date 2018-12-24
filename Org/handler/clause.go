package handler

import (
	"auditIntegralSys/Org/check"
	"auditIntegralSys/Org/db/clause"
	"auditIntegralSys/Org/entity"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Clause struct {
	gmvc.Controller
}

func (r *Clause) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.Clause
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	title := search.GetString("title")
	authorId := search.GetInt("authorId")
	state := search.GetString("state")

	searchMap := g.Map{}

	if title != "" {
		searchMap["title"] = title
	}

	if state != "" {
		searchMap["code"] = state
	}

	if authorId != 0 {
		searchMap["author_id"] = authorId
	}

	count, err := db_clause.GetClauseCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_clause.GetClauses(offset, size, searchMap)
		for _, v := range listData {
			rspData = append(rspData, entity.Clause{
				Id:           gconv.Int(v["id"]),
				DepartmentId: gconv.Int(v["department_id"]),
				Title:        gconv.String(v["title"]),
				AuthorId:     gconv.Int(v["author_id"]),
				AuthorName:   gconv.String(v["author_name"]),
				UpdateTime:   gconv.String(v["update_time"]),
				State:        gconv.String(v["state"]),
			})
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
	var err error = nil
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

func (r *Clause) Add() {
	reqData := r.Request.GetJson()
	departmentId := reqData.GetInt("departmentId")
	contentList := reqData.GetJson("content")
	cListLen := len(contentList.ToArray())
	state := reqData.GetString("state")

	id := 0
	msg := ""
	hasState := false
	hasDepartment := false
	var err error = nil
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
		id, err = db_clause.AddClause(g.Map{
			"department_id": departmentId,
			"title":         reqData.GetString("title"),
			"author_id":     util.GetUserIdByRequest(r.Cookie),
			"update_time":   util.GetLocalNowTimeStr(),
			"state":         state,
		})
	}
	// 添加内容
	if err == nil && id > 0 && cListLen > 0 {
		var contentArr []g.Map
		for i := 0; i < cListLen; i++ {
			contentArr = append(contentArr, g.Map{
				"clause_id":   id,
				"is_title":    contentList.GetBool(gconv.String(i) + ".isTitle"),
				"title_level": contentList.GetString(gconv.String(i) + ".titleLevel"),
				"content":     contentList.GetString(gconv.String(i) + ".content"),
				"order":       contentList.GetInt(gconv.String(i) + ".order"),
			})
		}
		_, err = db_clause.AddClauseContents(contentArr)
	}
	// 添加附件
	if err == nil && id > 0 {
		err = db_clause.AddClauseFiles(id, reqData.GetString("fileIds"))
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
			contentList = append(contentList, entity.ClauseContent{
				Id:         gconv.Int(v["id"]),
				ClauseId:   gconv.Int(v["clause_id"]),
				IsTitle:    gconv.Bool(v["is_title"]),
				TitleLevel: gconv.String(v["title_level"]),
				Content:    gconv.String(v["content"]),
				Order:      gconv.Int(v["order"]),
			})
		}
	}

	if clauseInfo.Id > 0 && err == nil {
		var fileRes []map[string]interface{}
		// 查询附件
		fileRes, err = db_clause.GetClauseFile(clauseInfo.Id)
		for _, v := range fileRes {
			fileList = append(fileList, entity2.File{
				Id:       gconv.Int(v["id"]),
				Name:     gconv.String(v["name"]),
				Suffix:   gconv.String(v["suffix"]),
				Time:     gconv.String(v["time"]),
				FileName: gconv.String(v["file_name"]),
				Path:     gconv.String(v["path"]),
			})
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
	var err error = nil
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
	var err error = nil
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
		_, _ = db_clause.DelClauseFile(clauseId)
		err = db_clause.AddClauseFiles(clauseId, reqData.GetString("fileIds"))
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
