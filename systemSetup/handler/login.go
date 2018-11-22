package handler

import (
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"auditIntegralSys/org/db/user"
	"auditIntegralSys/systemSetup/db/login"
	"auditIntegralSys/systemSetup/entity"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Login struct {
	gmvc.Controller
}

func (l *Login) List() {
	reqData := l.Request.GetJson()
	var rspData []*entity.User
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	userName := search.GetString("userName")
	key := search.GetString("key")
	departmentId := search.GetInt("departmentId")

	searchMap := g.Map{}

	if userName != "" {
		searchMap["user_name"] = userName
	}

	if key != "" {
		searchMap["'key'"] = key
	}

	if departmentId != 0 {
		searchMap["department_id"] = departmentId
	}

	count, err := db_login.GetUserCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_login.GetLoginList(offset, size, searchMap)
		for _, v := range listData {
			rspData = append(rspData, &entity.User{
				UserId:     gconv.Int(v["user_id"]),
				UserName:   gconv.String(v["user_name"]),
				AuthorName: gconv.String(v["author_name"]),
				LoginInfo: entity.LoginInfo{
					UserCode:     gconv.Int(v["user_code"]),
					IsUse:        gconv.Bool(v["is_use"]),
					LoginTime:    gconv.String(v["login_time"]),
					LoginNum:     gconv.Int(v["login_num"]),
					ChangePdTime: gconv.String(v["change_pd_time"]),
					AuthorId:     gconv.Int(v["author_id"]),
				},
			})
		}
	}
	if err != nil {
		log.Instance().Error(err)
	}
	l.Response.WriteJson(app.ListResponse{
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

func (l *Login) Add() {
	reqData := l.Request.GetJson()
	userCode := reqData.GetInt("userCode")
	isUse := reqData.GetBool("isUse")

	id := 0
	msg := ""
	hasCode := false
	// 检测人员表中是否存在此员工号
	checkCode, _, err := db_user.HasUserCode(userCode)
	if err == nil && checkCode {
		var userInfo entity.LoginInfo
		// 检测登录表中是否已经存在此员工号
		hasCode, userInfo, err = db_login.HasUserCode(userCode, true)
		if err == nil && (userInfo.Delete || !hasCode) {
			scStr := gconv.String(userCode)
			user := g.Map{
				"is_use":    gconv.Int(isUse),
				"author_id": 2,
				"password":  util.GetPasswordStr(scStr, scStr),
			}

			if hasCode {
				// 更新登录人员
				user["delete"] = 0
				_, err = db_login.UpdateLogin(user, userCode, 1)
				id = userInfo.LoginId
			} else {
				// 添加登录人员
				user["user_code"] = userCode
				id, err = db_login.AddLogin(user)
			}
		} else {
			msg = config.UserCode + config.Had
		}
	} else {
		msg = config.UserCode + config.NoHad
	}
	if err != nil {
		log.Instance().Error(err)
	}
	if msg == "" {
		msg = config.GetTodoResMsg(config.AddStr, err != nil)
	}
	l.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   msg,
		},
	})
}

func (l *Login) Edit() {
	reqData := l.Request.GetJson()
	userCode := reqData.GetInt("userCode")
	isUse :=  reqData.GetBool("isUse")
	rows, err := db_login.UpdateLogin(g.Map{"is_use": gconv.Int(isUse)}, userCode, 0)
	if err != nil {
		log.Instance().Error(err)
	}
	success := err == nil && rows > 0
	l.Response.WriteJson(app.Response{
		Data: userCode,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}

func (l *Login) Delete() {
	userCode := l.Request.GetQueryInt("userCode")
	rows, err := db_login.DelLogin(userCode)
	if err != nil {
		log.Instance().Error(err)
	}
	success := err == nil && rows > 0
	l.Response.WriteJson(app.Response{
		Data: userCode,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
