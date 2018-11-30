package handler

import (
	"auditIntegralSys/Org/db/department"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"

	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Department struct {
	gmvc.Controller
}

func (r *Department) List() {
	reqData := r.Request.GetJson()
	var rspData []entity.Department

	search := reqData.GetJson("search")
	title := search.GetString("title")
	parentId := search.GetInt("parentId")
	code := search.GetString("code")
	level := search.GetString("level")

	searchMap := g.Map{}

	if title != "" {
		searchMap["title"] = title
	}

	if code != "" {
		searchMap["code"] = code
	}
	if level != "" {
		searchMap["level"] = level
	}

	if parentId == 0 {
		// 默认根节点
		parentId = -1
	}

	listData, err := db_department.GetDepartmentsByParentId(parentId, searchMap)
	for _, v := range listData {
		rspData = append(rspData, entity.Department{
			Id:         gconv.Int(v["id"]),
			Name:       gconv.String(v["name"]),
			ParentId:   gconv.Int(v["parent_id"]),
			Code:       gconv.String(v["code"]),
			Level:      gconv.Int(v["level"]),
			Address:    gconv.String(v["address"]),
			Phone:      gconv.String(v["phone"]),
			UpdateTime: gconv.String(v["update_time"]),
		})
	}
	if err != nil {
		log.Instance().Errorfln("[Department List]: %v", err)
	}
	count := len(listData)
	r.Response.WriteJson(app.ListResponse{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.ListStr, err != nil),
		},
		Page: app.Pager{
			Page:  1,
			Size:  count,
			Total: count,
		},
	})
}

func (r *Department) Tree() {
	var rspData []entity.DepartmentTreeInfo
	parentId := r.Request.GetQueryInt("parentId")
	if parentId == 0 {
		// 默认根节点
		parentId = -1
	}

	listData, err := db_department.GetDepartmentsByParentId(parentId, g.Map{})
	for _, v := range listData {
		rspData = append(rspData, entity.DepartmentTreeInfo{
			Id:       gconv.Int(v["id"]),
			Name:     gconv.String(v["name"]),
			ParentId: gconv.Int(v["parent_id"]),
			Code:     gconv.String(v["code"]),
			Level:    gconv.Int(v["level"]),
		})
	}
	if err != nil {
		log.Instance().Errorfln("[Department Tree]: %v", err)
	}
	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.ListStr, err != nil),
		},
	})
}

func (r *Department) Add() {
	reqData := r.Request.GetJson()
	reqUserList := reqData.GetJson("userList")

	var userList []g.Map
	dictionaryType := g.Map{
		"parent_id":   reqData.GetInt("parentId"),
		"name":        reqData.GetString("name"),
		"code":        reqData.GetString("code"),
		"level":       reqData.GetInt("level"),
		"address":     reqData.GetString("address"),
		"phone":       reqData.GetString("phone"),
		"update_time": util.GetLocalNowTimeStr(),
	}

	id, err := db_department.AddDepartment(dictionaryType)
	userLen := len(reqUserList.ToArray())
	if err == nil && userLen > 0 {
		for i := 0; i < userLen; i++ {
			userList = append(userList, g.Map{
				"department_id": id,
				"user_id":       reqUserList.GetString(gconv.String(i) + ".userId"),
				"type":          reqUserList.GetString(gconv.String(i) + ".type"),
			})
		}
		_, err = db_department.AddDepartmentUser(userList)
		if err != nil {
			_, _ = db_department.DelDepartment(id)
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Department Add]: %v", err)
	}
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.AddStr, err != nil),
		},
	})
}

func (r *Department) Get() {
	departmentId := r.Request.GetQueryInt("id")
	userList := []entity.DepUser{}
	department, err := db_department.GetDepartment(departmentId)
	if err == nil && department.Id > 0 {
		var listData []map[string]interface{}
		listData, err = db_department.GetDepartmentUser(departmentId)
		for _, v := range listData {
			userList = append(userList, entity.DepUser{
				Id:       gconv.Int(v["id"]),
				UserId:   gconv.Int(v["user_id"]),
				UserName: gconv.String(v["user_name"]),
				Type:     gconv.String(v["type"]),
				TypeName: gconv.String(v["type_name"]),
			})
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Department Get]: %v", err)
	}
	success := err == nil && department.Id > 0
	r.Response.WriteJson(app.Response{
		Data: entity.DepartmentRes{
			Department: department,
			UserList:   userList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (r *Department) Edit() {
	reqData := r.Request.GetJson()
	departmentId := reqData.GetInt("id")
	userList := reqData.GetJson("userList")

	var addDepartments []g.Map
	var updateDepartments []g.Map
	var updateDepartmentUserIds []int
	department := g.Map{
		"parent_id":   reqData.GetInt("parentId"),
		"name":        reqData.GetString("name"),
		"code":        reqData.GetString("code"),
		"level":       reqData.GetInt("level"),
		"address":     reqData.GetString("address"),
		"phone":       reqData.GetString("phone"),
		"update_time": util.GetLocalNowTimeStr(),
	}

	rows, err := db_department.UpdateDepartment(departmentId, department)
	userLen := len(userList.ToArray())
	if err == nil && rows > 0 && userLen > 0 {
		for i := 0; i < userLen; i++ {
			id := userList.GetInt(gconv.String(i) + ".id")
			if id > 0 {
				updateDepartments = append(updateDepartments, g.Map{
					"id":            id,
					"department_id": departmentId,
					"user_id":       userList.GetString(gconv.String(i) + ".userId"),
					"type":          userList.GetString(gconv.String(i) + ".type"),
				})
				updateDepartmentUserIds = append(updateDepartmentUserIds, id)
			} else {
				addDepartments = append(addDepartments, g.Map{
					"department_id": departmentId,
					"user_id":       userList.GetString(gconv.String(i) + ".userId"),
					"type":          userList.GetString(gconv.String(i) + ".type"),
				})
			}
		}
		_, err = db_department.UpdateDepartmentUser(departmentId, addDepartments, updateDepartments, updateDepartmentUserIds)
	}
	if err != nil {
		log.Instance().Errorfln("[Department Edit]: %v", err)
	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: departmentId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}

func (r *Department) Delete() {
	departmentId := r.Request.GetQueryInt("id")
	rows, err := db_department.DelDepartment(departmentId)
	if err != nil {
		log.Instance().Errorfln("[Department Delete]: %v", err)

	}
	success := err == nil && rows > 0
	r.Response.WriteJson(app.Response{
		Data: departmentId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
