package handler

import (
	db_org_user "auditIntegralSys/Org/db/user"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/SystemSetup/db/login"
	ss_entity "auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/Worker/check"
	"auditIntegralSys/Worker/db/user"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/token"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type User struct {
	gmvc.Controller
}

func (u *User) Login() {
	reqData := u.Request.GetJson()
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
	} else {
		// 更新登录时间
		_, _ = db_login.UpdateLogin(g.Map{
			"login_time": util.GetLocalNowTimeStr(),
		}, userInfo.UserCode, 0)
	}
	success := err == nil && userId > 0 && checkPd
	if msg == "" {
		msg = config.GetTodoResMsg(config.LoginStr, !success)
	}
	if success {
		token.Set(userId, u.Request, true)
	} else {
		token.Del(u.Request)
	}
	u.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (u *User) Get() {
	userId := util.GetUserIdByRequest(u.Cookie)
	log.Instance().Infofln("userId %v", userId)

	userInfo, err := db_org_user.GetUser(userId)
	if err != nil {
		log.Instance().Errorfln("[User Get]: %v", err)
	}
	success := err == nil && userInfo.UserId > 0
	if success {
		token.Set(userId, u.Request, false)
	}
	u.Response.WriteJson(app.Response{
		Data: userInfo,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr+config.UserInfoStr, !success),
		},
	})
}

func (u *User) Password() {
	msg := ""
	reqData := u.Request.GetJson()
	oldPd := reqData.GetString("password")
	newPd := reqData.GetString("new")

	rows := 0
	userId := util.GetUserIdByRequest(u.Cookie)

	var err error = nil
	var userInfo ss_entity.LoginInfo

	if !check.PasswordLen(newPd) {
		msg = config.PasswordLenErrStr
		err = errors.New(msg)
	}
	if err == nil {
		userInfo, err = db_login.GetLoginUserInfoByUserId(userId)
	}
	if err == nil && userInfo.UserId != 0 && check.Password(userInfo.UserCode, oldPd, userInfo.Password) {
		rows, err = db_login.UpdateLogin(g.Map{
			"change_pd_time": util.GetLocalNowTimeStr(),
			"password":       util.GetPasswordStr(newPd, gconv.String(userInfo.UserCode)),
		}, userInfo.UserCode, 0)
	} else if msg == "" {
		msg = config.PasswordErrStr
	}
	if err != nil {
		log.Instance().Errorfln("[User Password]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.ChangePasswordStr, !success)
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

func (u *User) Logout() {
	token.Del(u.Request)
	u.Response.WriteJson(app.Response{
		Data: "",
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   config.GetTodoResMsg(config.LogoutStr, false),
		},
	})
}
