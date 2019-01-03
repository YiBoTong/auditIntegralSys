package handler

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/rectifyReport"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Worker/db/file"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/encoding/gjson"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type RectifyReport struct {
	gmvc.Controller
}

func (r *RectifyReport) checkId(id int) (bool, string) {
	msg := ""
	canEdit := true
	if id == 0 { // 状态和ID都必须要有
		canEdit = false
	}
	return canEdit, msg
}

// 检测状态是否合法
func (r *RectifyReport) checkState(state string) (bool, string) {
	hasState, msg := check.PublicState(state).Has()
	return hasState, msg
}

func (r *RectifyReport) checkIdAndState(id int, state string) (bool, string) {
	canEdit, msg := r.checkId(id)
	if canEdit {
		canEdit, msg = r.checkState(state)
	}
	return canEdit, msg
}

func (r *RectifyReport) addCall(rectifyId, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	fileIds := json.GetString("fileIds")
	contentList := json.GetJson("contentList")
	ContentList := []g.Map{}
	data := g.Map{
		"state":       stateStr,
		"rectify_id":  rectifyId,
		"update_time": util.GetLocalNowTimeStr(),
	}
	contentListItem := g.Map{
		"draft_content_id": "int",
		"content":          "string",
	}
	util.GetSqlMapItemFun(*contentList, contentListItem, func(itemMap g.Map) {
		ContentList = append(ContentList, itemMap)
	})
	row, err := db_rectifyReport.Add(rectifyId, fileIds, data, ContentList)
	return row, err
}

func (r *RectifyReport) editCall(id, todoUserId int, stateStr string, json gjson.Json) (int, error) {
	fileIds := json.GetString("fileIds")
	contentList := json.GetJson("contentList")
	ContentList := []g.Map{}
	data := g.Map{
		"state":       stateStr,
		"update_time": util.GetLocalNowTimeStr(),
	}
	contentListItem := g.Map{
		"draft_content_id": "int",
		"content":          "string",
	}
	util.GetSqlMapItemFun(*contentList, contentListItem, func(itemMap g.Map) {
		ContentList = append(ContentList, itemMap)
	})
	row, err := db_rectifyReport.Edit(id, fileIds, data, ContentList)
	return row, err
}

func (r *RectifyReport) Get() {
	var err error = nil
	id := r.Request.GetQueryInt("id")
	rectifyId := r.Request.GetQueryInt("rectifyId")
	ContentList := []entity.RectifyReportContentItem{}
	FileList := []entity2.File{}
	RectifyReportItem := entity.RectifyReportItem{}

	if id != 0 {
		RectifyReportItem, err = db_rectifyReport.Get(id)
	} else if rectifyId != 0 {
		RectifyReportItem, err = db_rectifyReport.GetByRectifyId(rectifyId)
	}

	if err != nil {
		log.Instance().Errorfln("[Rectify Get]: %v", err)
	}

	if RectifyReportItem.Id != 0 {
		contentList, _ := db_rectifyReport.GetContents(RectifyReportItem.Id)
		fileList, _ := db_file.GetFilesByFrom(RectifyReportItem.Id,table.RectifyReport)
		for _, bv := range contentList {
			item := entity.RectifyReportContentItem{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				ContentList = append(ContentList, item)
			}
		}
		for _, bv := range fileList {
			item := entity2.File{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				FileList = append(FileList, item)
			}
		}
	}

	success := err == nil && RectifyReportItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.RectifyReport{
			RectifyReportItem: RectifyReportItem,
			ContentList:       ContentList,
			FileList:          FileList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *RectifyReport) Edit() {
	rows := 0
	var err error = nil
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	rectifyId := reqData.GetInt("rectifyId")
	stateStr := reqData.GetString("state")
	todoUserId := util.GetUserIdByRequest(r.Request.Cookie)
	checkRes, msg := r.checkState(stateStr)
	if checkRes {
		if id != 0 {
			rows, err = r.editCall(id, todoUserId, stateStr, *reqData)
		} else if rectifyId != 0 {
			rows, err = r.addCall(rectifyId, todoUserId, stateStr, *reqData)
		}
	}
	if err != nil {
		log.Instance().Errorfln("[PunishNotice Edit]: %v", err)
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
