package handler

import (
	"auditIntegralSys/Worker/db/menu"
	"auditIntegralSys/Worker/fun"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
)

type Menu struct {
	gmvc.Controller
}

func (r *Menu) Get() {
	allMenu, err := fun.GetRbacMenu(-1, "management")
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
