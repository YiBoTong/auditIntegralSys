package handler

import (
	"auditIntegralSys/Org/db/department"
	db_org_user "auditIntegralSys/Org/db/user"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/SystemSetup/db/login"
	ss_entity "auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/Worker/check"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/Worker/db/rbac"
	"auditIntegralSys/Worker/db/user"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/table"
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
	userCode := reqData.GetString("userCode")
	password := reqData.GetString("password")
	msg := ""
	checkPd := false
	userId := 0
	userInfo := entity.User{}
	loginUser := ss_entity.LoginInfo{}
	portraitFile := entity2.File{}

	Departments := []entity.LoginUserDepartmentItem{}
	RbacList := []entity2.RbacListItem{}

	err := error(nil)
	if password == "" {
		msg = "密码不能为空"
	} else if userCode == "" {
		msg = "员工号不能为空"
	}
	if msg == "" {
		checkPd, loginUser, err = db_user.Login(userCode, password)
		userId = loginUser.UserId
		if userId == 0 {
			msg = "该员工号不允许登录"
		} else if !checkPd {
			msg = "密码不正确"
		}
	}
	if msg == "" && err == nil {
		userRbacs := g.Slice{}
		userInfo, err = db_org_user.GetUser(userId)
		portraitFile, _ = db_file.Get(userInfo.PortraitId)
		departmentList, _ := db_department.GetUserDepartmentByUserId(userId)
		for _, v := range departmentList {
			item := entity.LoginUserDepartmentItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				userRbacs = append(userRbacs, item.Type)
				Departments = append(Departments, item)
			}
		}
		if len(userRbacs) >0 {
			rbacList, _ := db_rbac.GetUserRbacByKeys(userRbacs)
			for _, rv := range rbacList {
				item := entity2.RbacListItem{}
				if ok := gconv.Struct(rv, &item); ok == nil {
					RbacList = append(RbacList, item)
				}
			}
		}else{
			msg = "当前员工号无任何角色，不能进入系统"
		}
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
			"login_num":  loginUser.LoginNum + 1,
		}, userInfo.UserCode, 0)
	}
	success := err == nil && userId != 0 && checkPd
	if msg == "" {
		msg = config.GetTodoResMsg(config.LoginStr, !success)
	}
	if success {
		token.Set(userId, u.Request, true)
	} else {
		token.Del(u.Request)
	}
	u.Response.WriteJson(app.Response{
		Data: entity.LoginUserInfo{
			User:         userInfo,
			RbacList:     RbacList,
			Departments:  Departments,
			PortraitFile: portraitFile,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (u *User) Get() {
	userId := util.GetUserIdByRequest(u.Cookie)

	userInfo, err := db_org_user.GetUser(userId)
	portraitFile, _ := db_file.Get(userInfo.PortraitId)
	departmentList, _ := db_department.GetUserDepartmentByUserId(userId)

	userRbacs := g.Slice{}
	Departments := []entity.LoginUserDepartmentItem{}
	RbacList := []entity2.RbacListItem{}

	for _, v := range departmentList {
		item := entity.LoginUserDepartmentItem{}
		if ok := gconv.Struct(v, &item); ok == nil {
			userRbacs = append(userRbacs, item.Type)
			Departments = append(Departments, item)
		}
	}

	rbacList, _ := db_rbac.GetUserRbacByKeys(userRbacs)
	for _, rv := range rbacList {
		item := entity2.RbacListItem{}
		if ok := gconv.Struct(rv, &item); ok == nil {
			RbacList = append(RbacList, item)
		}
	}

	if err != nil {
		log.Instance().Errorfln("[User Get]: %v", err)
	}
	success := err == nil && userInfo.UserId != 0
	if success {
		token.Set(userId, u.Request, false)
	}
	u.Response.WriteJson(app.Response{
		Data: entity.LoginUserInfo{
			User:         userInfo,
			RbacList:     RbacList,
			Departments:  Departments,
			PortraitFile: portraitFile,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr+config.UserInfoStr, !success),
		},
	})
}

func (r *User) Edit() {
	reqData := r.Request.GetJson()
	userId := util.GetUserIdByRequest(r.Cookie)

	Update := g.Map{}
	update := g.Map{
		"phone":       "string",
		"portrait_id": "int",
	}

	for k, v := range update {
		util.GetSearchMapByReqJson(Update, *reqData, k, gconv.String(v))
	}

	rows, err := db_org_user.UpdateUser(userId, Update)
	if err == nil {
		if Update["portrait_id"] != 0 {
			_, err = db_file.UpdateFileByIds(table.User, gconv.String(Update["portrait_id"]), userId)
		} else {
			_, err = db_file.DelFilesByFrom(userId, table.User)
		}
	}

	if err != nil {
		log.Instance().Errorfln("[Clause Edit]: %v", err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: userId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
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

	err := error(nil)
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
