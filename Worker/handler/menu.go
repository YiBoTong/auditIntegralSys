package handler

import (
	"auditIntegralSys/Org/db/department"
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/menu"
	"auditIntegralSys/Worker/fun"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Menu struct {
	gmvc.Controller
}

func (r *Menu) Get() {
	userId := util.GetUserIdByRequest(r.Cookie)
	userRbacs := g.Slice{}

	departmentList, _ := db_department.GetUserDepartmentByUserId(userId)
	for _, v := range departmentList {
		item := entity.LoginUserDepartmentItem{}
		if ok := gconv.Struct(v, &item); ok == nil {
			userRbacs = append(userRbacs, item.Type)
		}
	}

	allMenu, err := fun.GetRbacMenu(-1, userRbacs)
	if err != nil {
		allMenu = nil
		log.Instance().Errorfln("[Worker Menu Get]: %v", err)
	}
	r.Response.WriteJson(app.Response{
		Data: allMenu,
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   config.GetTodoResMsg(config.GetStr, false),
		},
	})
}

func (r *Menu) All() {
	allMenu, err := fun.GetAllMenu(-1, true)
	if err != nil {
		allMenu = nil
		log.Instance().Errorfln("[Worker Menu All]: %v", err)
	}
	r.Response.WriteJson(app.Response{
		Data: allMenu,
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   config.GetTodoResMsg(config.GetStr, false),
		},
	})
}

func (r *Menu) Add() {
	reqData := r.Request.GetJson()
	msg := ""
	id, err := db_menu.Add(g.Map{
		"path":      reqData.GetString("path"),
		"icon":      reqData.GetString("icon"),
		"title":     reqData.GetString("title"),
		"name":      reqData.GetString("name"),
		"no_cache":  reqData.GetString("noCache"),
		"order":     reqData.GetString("Order"),
		"parent_id": reqData.GetString("parentId"),
		"time":      util.GetLocalNowTimeStr(),
	})
	success := err == nil && id > 0
	if err != nil {
		log.Instance().Errorfln("[Worker Menu Add]: %v", err)
	}
	if msg == "" {
		msg = config.GetTodoResMsg(config.AddStr, !success)
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
