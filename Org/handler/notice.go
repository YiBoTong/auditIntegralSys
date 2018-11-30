package handler

import (
	"auditIntegralSys/Org/check"
	"auditIntegralSys/Org/db/notice"
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
	"strings"
)

type Notice struct {
	gmvc.Controller
}

func (r *Notice) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.NoticeList
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	title := search.GetString("title")
	state := search.GetString("state")

	searchMap := g.Map{}

	if title != "" {
		searchMap["title"] = title
	}

	if state != "" {
		searchMap["state"] = state
	}

	count, err := db_notice.GetNoticeCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_notice.GetNotices(offset, size, searchMap)
		for _, v := range listData {
			rspData = append(rspData, entity.NoticeList{
				Id:           gconv.Int(v["id"]),
				DepartmentId: gconv.Int(v["department_id"]),
				Title:        gconv.String(v["title"]),
				Time:         gconv.String(v["time"]),
				Range:        gconv.Int(v["range"]),
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

func (r *Notice) Add() {
	reqData := r.Request.GetJson()
	departmentId := reqData.GetInt("departmentId")
	rangeType := reqData.GetInt("range")
	state := reqData.GetString("state")

	id := 0
	hasState := false
	// 检测是否部门是否存在
	hasDepartment, msg, err := check.HasDepartment(departmentId)
	if !hasDepartment {
		err = errors.New(msg)
	}
	if err == nil {
		// 检测状态是否合法
		hasState, msg = check.NoticeState(state).HasState()
		if !hasState {
			err = errors.New(msg)
		}
	}
	// 添加通知
	if hasDepartment && err == nil {
		id, err = db_notice.AddNotice(g.Map{
			"department_id": departmentId,
			"title":         reqData.GetString("title"),
			"content":       reqData.GetString("content"),
			"time":          util.GetLocalNowTimeStr(),
			"range":         rangeType,
			"state":         state,
		})
	}
	// 添加指定的通知部门
	if rangeType == 2 && id > 0 {
		msg, err = addNoticeInform(id, reqData.GetString("informIds"))
	}
	// 添加附件
	if err == nil && id > 0 {
		err = db_notice.AddNoticeFiles(id, reqData.GetString("fileIds"))
	}
	if err != nil && id > 0 {
		_, _ = db_notice.DelNotice(id)
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

func (r *Notice) Get() {
	noticeId := r.Request.GetQueryInt("id")
	fileList := []entity2.File{}
	informs := []entity.DepartmentTreeInfo{}
	noticeInfo, err := db_notice.GetNotice(noticeId)
	if noticeInfo.Range == 2 {
		var dep []map[string]interface{}
		// 获取被通知的部门列表
		dep, err = db_notice.GetNoticeInform(noticeId)
		for _, v := range dep {
			informs = append(informs, entity.DepartmentTreeInfo{
				Id:       gconv.Int(v["id"]),
				Name:     gconv.String(v["name"]),
				ParentId: gconv.Int(v["parent_id"]),
				Code:     gconv.String(v["code"]),
				Level:    gconv.Int(v["level"]),
			})
		}
	}
	if noticeInfo.Id > 0 && err == nil {
		var fileRes []map[string]interface{}
		// 查询附件
		fileRes, err = db_notice.GetNoticeFile(noticeId)
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
	success := err == nil && noticeInfo.Id > 0
	r.Response.WriteJson(app.Response{
		Data: entity.NoticeRes{
			Notice:   noticeInfo,
			FileList: fileList,
			Informs:  informs,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Notice) Edit() {
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	departmentId := reqData.GetInt("departmentId")
	rangeType := reqData.GetInt("range")
	state := reqData.GetString("state")

	rows := 0
	hasState := false
	// 检测是否部门是否存在
	hasDepartment, msg, err := check.HasDepartment(departmentId)
	if !hasDepartment {
		err = errors.New(msg)
	}
	if err == nil {
		// 检测状态是否合法
		hasState, msg = check.NoticeState(state).HasState()
		if !hasState {
			err = errors.New(msg)
		}
	}
	// 添加通知
	if hasDepartment && err == nil {
		rows, err = db_notice.UpdateNotice(id, g.Map{
			"department_id": departmentId,
			"title":         reqData.GetString("title"),
			"content":       reqData.GetString("content"),
			"time":          util.GetLocalNowTimeStr(),
			"range":         rangeType,
			"state":         reqData.GetString("state"),
		})
	}
	// 添加指定的通知部门
	if rangeType == 2 {
		_, _ = db_notice.DelNoticeInform(id)
		msg, err = addNoticeInform(id, reqData.GetString("informIds"))
	}
	// 添加附件
	if err == nil && rows > 0 {
		_, _ = db_notice.DelNoticeFile(id)
		err = db_notice.AddNoticeFiles(id, reqData.GetString("fileIds"))
	}
	if err != nil && rows > 0 {
		log.Instance().Errorfln("[Notice Edit]: %v", err)
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

func (r *Notice) State() {
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	state := reqData.GetString("state")
	rows := 0
	var err error = nil
	// 检测状态是否合法
	hasState, msg := check.NoticeState(state).HasState()
	if hasState {
		rows, err = db_notice.UpdateNotice(id, g.Map{
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

func (r *Notice) Delete() {
	noticeId := r.Request.GetQueryInt("id")
	rows, err := db_notice.DelNotice(noticeId)
	if err != nil {
		log.Instance().Error(err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: noticeId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}

func addNoticeInform(noticeId int, informIds string) (string, error) {
	msg := ""
	var err error = nil
	ids := strings.Split(informIds, ",")
	if len(ids) > 0 && ids[0] != "" {
		var add []g.Map
		for _, id := range ids {
			dId := gconv.Int(id)
			if dId > 0 {
				add = append(add, g.Map{
					"notice_id":     noticeId,
					"department_id": dId,
				})
			}
		}
		_, err = db_notice.AddNoticeInform(add)
	} else {
		msg = "指定的部门不能为空"
		err = errors.New(msg)
	}
	return msg, err
}
