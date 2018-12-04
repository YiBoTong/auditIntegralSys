package handler

import (
	"auditIntegralSys/Worker/db/menu"
	"auditIntegralSys/Worker/entity"
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

func getMenus(parentId int) []entity.Menus {
	var allMenu []entity.Menus
	menuList, err := db_menu.Menus(parentId)
	for _, v := range menuList {
		allMenu = append(allMenu, entity.Menus{
			Menu: entity.Menu{
				Id:       gconv.Int(v["id"]),
				Path:     gconv.String(v["path"]),
				Title:    gconv.String(v["title"]),
				Icon:     gconv.String(v["icon"]),
				NoCache:  gconv.Bool(v["no_cache"]),
				Order:    gconv.Int(v["order"]),
				ParentId: gconv.Int(v["parentId"]),
				HasChild: gconv.Bool(v["has_child"]),
				Time:     gconv.String(v["time"]),
			},
			Children: nil,
		})
	}
	if err != nil {
		log.Instance().Errorfln("[All Menu]: %v", err)
	}
	return allMenu
}

func (r *Menu) All() {
	var allMenu []entity.Menus
	menuList, err := db_menu.Menus(-1)
	for _, v := range menuList {
		allMenu = append(allMenu, entity.Menus{
			Menu: entity.Menu{
				Id:       gconv.Int(v["id"]),
				Path:     gconv.String(v["path"]),
				Title:    gconv.String(v["title"]),
				Icon:     gconv.String(v["icon"]),
				NoCache:  gconv.Bool(v["no_cache"]),
				Order:    gconv.Int(v["order"]),
				ParentId: gconv.Int(v["parentId"]),
				HasChild: gconv.Bool(v["has_child"]),
				Time:     gconv.String(v["time"]),
			},
			Children: nil,
		})
	}
	if err != nil {
		log.Instance().Errorfln("[All Menu]: %v", err)
	}
	log.Instance().Debugfln("all %v",allMenu)
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
