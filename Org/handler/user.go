package handler

import (
	"auditIntegralSys/Org/check"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Org/fun"
	"auditIntegralSys/Worker/db/file"
	entity2 "auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"regexp"
)

type User struct {
	gmvc.Controller
}

var xlsxRegexp = regexp.MustCompile(`^xlsx$`)

func (r *User) importCall(departmentId int, xlsxRows [][]string) (g.List, g.Slice) {
	addUserCode := g.Slice{}
	addUser := g.List{}
	sexKeys := g.Map{"男": 2, "女": 1}
	userKeys := []string{
		"user_name",
		"user_code",
		"class",
		"sex",
		"phone",
		"id_card",
	}
	updateTime := util.GetLocalNowTimeStr()
	for i, row := range xlsxRows {
		if i == 0 {
			continue
		}
		userItem := g.Map{}
		for j, colCell := range row {
			if j == 1 {
				addUserCode = append(addUserCode, colCell)
			}
			userItem[userKeys[j]] = colCell
		}
		sex := 0
		sexVal := sexKeys[gconv.String(userItem["sex"])]
		if sexVal != "" {
			userItem["sex"] = sexVal
		} else {
			userItem["sex"] = sex
		}
		userItem["update_time"] = updateTime
		userItem["department_id"] = departmentId
		addUser = append(addUser, userItem)
	}
	return addUser, addUserCode
}

func (r *User) List() {
	reqData := r.Request.GetJson()
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
		"user_name":                     "string",
		"u.user_code:user_code":         "int",
		"u.department_id:department_id": "int",
		"u.sex:sex":                     "int",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
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

func (r *User) Import() {
	fileId := r.Request.GetQueryInt("fileId")
	departmentId := r.Request.GetQueryInt("departmentId")
	id := 0
	msg := ""
	errorCode := 0
	file := entity2.File{}
	xlsxRows := [][]string{}
	HadUserList := []entity.User{}
	var err error
	if departmentId == 0 || departmentId == -1 {
		departmentId = -1
	} else {
		// 检测是否部门是否存在
		hasDepartment := false
		hasDepartment, msg, err = check.HasDepartment(departmentId)
		if !hasDepartment {
			err = errors.New(msg)
		}
	}
	if err == nil {
		file, err = db_file.Get(fileId)
	}
	if err == nil && file.Id != 0 {
		filePath := g.Config().GetString("filePath") + file.Path + file.FileName + "." + file.Suffix
		if xlsxRegexp.MatchString(file.Suffix) {
			xlsxRows, err = fun.ReadSpreadsheets(filePath, "人员导入模版")
		}
	}
	if err == nil {
		if len(xlsxRows) > 1 {
			addUsers, addUserCodes := r.importCall(departmentId, xlsxRows)
			if len(addUsers) > 0 {
				hasUserList := g.List{}
				hasUserList, err = db_user.HasUserCodes(addUserCodes)
				if len(hasUserList) == 0 && err == nil {
					id, err = db_user.AddUser(addUsers)
				} else {
					for _, v := range hasUserList {
						item := entity.User{}
						if ok := gconv.Struct(v, &item); ok == nil {
							HadUserList = append(HadUserList, item)
						}
					}
					errorCode = 2
				}
			}
		}
	}
	success := err == nil && id > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.ImportStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: entity.ImportUserRes{
			Id:          id,
			HadUserList: HadUserList,
		},
		Status: app.Status{
			Code:  errorCode,
			Error: !success,
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
