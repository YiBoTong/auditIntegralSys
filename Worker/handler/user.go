package handler

import (
	db_org_user "auditIntegralSys/Org/db/user"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/user"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/token"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g/frame/gmvc"
)

type User struct {
	gmvc.Controller
}

func (l *User) Login() {
	reqData := l.Request.GetJson()
	userCode := reqData.GetInt("userCode")
	password := reqData.GetString("password")
	msg := ""
	checkPd := false
	userId := 0
	var err error = nil
	var userInfo entity.User
	if password == "" {
		msg = "密码不能为空"
	} else if userCode == 0 {
		msg = "员工号不能为空"
	}
	if msg == "" {
		checkPd, userId, err = db_user.Login(userCode, password)
		if userId == 0 {
			msg = "该员工号不允许登录"
		} else if !checkPd {
			msg = "密码不正确"
		}
	}
	if msg == "" && err == nil {
		userInfo, err = db_org_user.GetUser(userId)
	}
	if msg != "" {
		err = errors.New(msg)
	}
	if err != nil {
		log.Instance().Errorfln("[User Login]: %v", err)
	}
	success := err == nil && userId > 0 && checkPd
	if msg == "" {
		msg = config.GetTodoResMsg(config.LoginStr, !success)
	}
	if success {
		token.Set(userId, l.Request, true)
	} else {
		token.Del(l.Request)
	}
	l.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (l *User) Get() {
	userId := util.GetUserIdByRequest(l.Cookie)
	log.Instance().Infofln("userId %v", userId)

	userInfo, err := db_org_user.GetUser(userId)
	success := err == nil && userInfo.UserId > 0
	if success {
		token.Set(userId, l.Request, false)
	}
	l.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr+config.UserInfoStr, !success),
		},
	})
}

func (l *User) Logout() {
	token.Del(l.Request)
	l.Response.WriteJson(app.Response{
		Data: "",
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   config.GetTodoResMsg(config.LogoutStr, false),
		},
	})
}
