package handler

import (
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/SystemSetup/db/login"
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Login struct {
	gmvc.Controller
}

func (l *Login) List() {
	reqData := l.Request.GetJson()
	rspData := []entity.User{}
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"user_name":             "string",
		"u.user_code:user_code": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// d.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	count, err := db_login.GetUserCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_login.GetLoginList(offset, size, listSearchMap)
		for _, v := range listData {
			LoginUser := entity.LoginUser{}
			LoginInfo := entity.LoginInfo{}
			o := gconv.Struct(v, &LoginUser)
			k := gconv.Struct(v, &LoginInfo)
			if o == nil && k == nil {
				item := entity.User{
					LoginUser: LoginUser,
					LoginInfo: LoginInfo,
				}
				rspData = append(rspData, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Login List]: %v", err)
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
	userCode := reqData.GetString("userCode")
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
		log.Instance().Errorfln("[Login Add]: %v", err)
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
	userCode := reqData.GetString("userCode")
	isUse := reqData.GetBool("isUse")
	rows, err := db_login.UpdateLogin(g.Map{"is_use": gconv.Int(isUse)}, userCode, 0)
	if err != nil {
		log.Instance().Errorfln("[Login Edit]: %v", err)
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
	userCode := l.Request.GetQueryString("userCode")
	rows, err := db_login.DelLogin(userCode)
	if err != nil {
		log.Instance().Errorfln("[Login Delete]: %v", err)
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
