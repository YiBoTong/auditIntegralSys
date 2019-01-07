package handler

import (
	"auditIntegralSys/Org/check"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type User struct {
	gmvc.Controller
}

func (r *User) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.User
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"user_name":     "string",
		"user_code":     "int",
		"department_id": "int",
		"sex":           "int",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "u."+k+":"+k, gconv.String(v))
	}

	count, err := db_user.GetUserCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_user.GetUsers(offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.User{}
			err = gconv.Struct(v, &item)
			if err == nil {
				rspData = append(rspData, item)
			} else {
				break
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[User List]: %v", err)
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

func (r *User) Add() {
	reqData := r.Request.GetJson()
	userCode := reqData.GetString("userCode")
	departmentId := reqData.GetInt("departmentId")

	id := 0
	hasDepartment := false
	// 检测人员表中是否存在此员工号
	hasCode, msg, err := check.HasUserCode(userCode, 0)
	if !hasCode && err == nil {
		// 检测是否部门是否存在
		hasDepartment, msg, err = check.HasDepartment(departmentId)
		if !hasDepartment {
			err = errors.New(msg)
		}
	}
	if hasDepartment && err == nil {
		var user []g.Map
		user = append(user, g.Map{
			"department_id": departmentId,
			"user_code":     userCode,
			"user_name":     reqData.GetString("userName"),
			"phone":         reqData.GetString("phone"),
			"sex":           reqData.GetInt("sex"),
			"class":         reqData.GetString("class"),
			"id_card":       reqData.GetString("idCard"),
			"update_time":   util.GetLocalNowTimeStr(),
		})
		id, err = db_user.AddUser(user)
	}

	if err != nil {
		log.Instance().Errorfln("[Login Add]: %v", err)
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

func (r *User) Get() {
	userId := r.Request.GetQueryInt("id")
	userInfo, err := db_user.GetUser(userId)
	if err != nil {
		log.Instance().Errorfln("[User Get]: %v", err)
	}
	success := err == nil && userInfo.UserId != 0
	r.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *User) Edit() {
	reqData := r.Request.GetJson()
	userId := reqData.GetInt("userId")
	userCode := reqData.GetString("userCode")
	departmentId := reqData.GetInt("departmentId")

	rows := 0
	hasDepartment := false
	// 检测人员表中是否存在此员工号
	hasCode, msg, err := check.HasUserCode(userCode, userId)
	if !hasCode && err == nil {
		// 检测是否部门是否存在
		hasDepartment, msg, err = check.HasDepartment(departmentId)
		if !hasDepartment {
			err = errors.New(msg)
		}
	}
	if hasDepartment && err == nil {
		user := g.Map{
			"department_id": departmentId,
			"user_code":     userCode,
			"user_name":     reqData.GetString("userName"),
			"phone":         reqData.GetString("phone"),
			"sex":           reqData.GetInt("sex"),
			"class":         reqData.GetString("class"),
			"id_card":       reqData.GetString("idCard"),
			"update_time":   util.GetLocalNowTimeStr(),
		}
		rows, err = db_user.UpdateUser(userId, user)
	}

	if err != nil {
		log.Instance().Errorfln("[Login Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.EditStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: userId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (r *User) Delete() {
	userId := r.Request.GetQueryInt("id")
	rows, err := db_user.DelUser(userId)
	if err != nil {
		log.Instance().Error(err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: userId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
