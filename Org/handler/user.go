package handler

import (
	"auditIntegralSys/Org/db/department"
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

func (u *User) List() {
	reqData := u.Request.GetJson()
	var rspData []entity.User
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	userName := search.GetString("userName")
	userCode := search.GetInt("userCode")
	departmentId := search.GetInt("departmentId")
	sex := search.GetInt("sex")

	searchMap := g.Map{}

	if userName != "" {
		searchMap["user_name"] = userName
	}

	if userCode != 0 {
		searchMap["user_code"] = userCode
	}

	if departmentId != 0 {
		searchMap["department_id"] = departmentId
	}

	if sex != 0 {
		searchMap["sex"] = sex
	}

	count, err := db_user.GetUserCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_user.GetUsers(offset, size, searchMap)
		for _, v := range listData {
			rspData = append(rspData, entity.User{
				UserId:       gconv.Int(v["user_id"]),
				DepartmentId: gconv.Int(v["department_id"]),
				UserName:     gconv.String(v["user_name"]),
				UserCode:     gconv.Int(v["user_code"]),
				Class:        gconv.String(v["class"]),
				Phone:        gconv.String(v["phone"]),
				IdCard:       gconv.String(v["id_card"]),
				UpdateTime:   gconv.String(v["update_time"]),
			})
		}
	}
	if err != nil {
		log.Instance().Errorf("[User List]: %v", err)
	}
	u.Response.WriteJson(app.ListResponse{
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

func (u *User) Add() {
	reqData := u.Request.GetJson()
	userCode := reqData.GetInt("userCode")
	departmentId := reqData.GetInt("departmentId")

	id := 0
	hasDepartment := false
	// 检测人员表中是否存在此员工号
	hasCode, msg, err := checkHasUserCode(userCode, 0)
	if !hasCode && err == nil {
		// 检测是否部门是否存在
		hasDepartment, msg, err = checkHasDepartment(departmentId)
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
			"sex":           reqData.GetInt("sex"),
			"class":         reqData.GetString("class"),
			"id_card":       reqData.GetString("idCard"),
			"update_time":   util.GetLocalNowTimeStr(),
		})
		id, err = db_user.AddUser(user)
	}

	if err != nil {
		log.Instance().Errorf("[Login Add]: %v", err)
	}
	if msg == "" {
		msg = config.GetTodoResMsg(config.AddStr, err != nil)
	}
	u.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   msg,
		},
	})
}

func (u *User) Get() {
	userId := u.Request.GetQueryInt("id")
	userInfo, err := db_user.GetUser(userId)
	if err != nil {
		log.Instance().Errorf("[User Get]: %v", err)
	}
	success := err == nil && userInfo.UserId > 0
	u.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (u *User) Edit() {
	reqData := u.Request.GetJson()
	userId := reqData.GetInt("userId")
	userCode := reqData.GetInt("userCode")
	departmentId := reqData.GetInt("departmentId")

	rows := 0
	hasDepartment := false
	// 检测人员表中是否存在此员工号
	hasCode, msg, err := checkHasUserCode(userCode, userId)
	if !hasCode && err == nil {
		// 检测是否部门是否存在
		hasDepartment, msg, err = checkHasDepartment(departmentId)
		if !hasDepartment {
			err = errors.New(msg)
		}
	}
	if hasDepartment && err == nil {
		user := g.Map{
			"department_id": departmentId,
			"user_code":     userCode,
			"user_name":     reqData.GetString("userName"),
			"sex":           reqData.GetInt("sex"),
			"class":         reqData.GetString("class"),
			"id_card":       reqData.GetString("idCard"),
			"update_time":   util.GetLocalNowTimeStr(),
		}
		rows, err = db_user.UpdateUser(userId, user)
	}

	if err != nil {
		log.Instance().Errorf("[Login Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.EditStr, !success)
	}
	u.Response.WriteJson(app.Response{
		Data: userId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (u *User) Delete() {
	userId := u.Request.GetQueryInt("id")
	rows, err := db_user.DelUser(userId)
	if err != nil {
		log.Instance().Error(err)
	}
	success := err == nil && rows > 0
	u.Response.WriteJson(app.Response{
		Data: userId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}


// 检测员工好是否存在（传userId将排除此userId后检测）
func checkHasUserCode(userCode int, userId int) (bool, string, error) {
	msg := ""
	hasCode, userInfo, err := db_user.HasUserCode(userCode)
	if userId > 0 && userInfo.UserId == userId {
		hasCode = false
	} else if err == nil {
		msg = config.UserCode
		if hasCode {
			msg += config.Had
		} else {
			msg += config.NoHad
		}
	}
	return hasCode, msg, err
}

func checkHasDepartment(departmentId int) (bool, string, error) {
	msg := ""
	hasDepartment, err := db_department.HasDepartment(departmentId)
	if err == nil && !hasDepartment {
		msg = config.DepartmentMsgStr + config.NoHad
	}
	return hasDepartment, msg, err
}
