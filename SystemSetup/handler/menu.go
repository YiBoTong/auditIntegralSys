package handler

import (
	"auditIntegralSys/Worker/db/menu"
	"auditIntegralSys/Worker/fun"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Menu struct {
	gmvc.Controller
}

func (r *Menu) All() {
	allMenu, err := fun.GetAllMenu(-1, false)
	if err != nil {
		allMenu = nil
		log.Instance().Errorfln("[Menu All]: %v", err)
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

func (r *Menu) IsUse() {
	reqData := r.Request.GetJson()
	id := reqData.GetInt("id")
	row, err := db_menu.Update(id, g.Map{
		"is_use": gconv.Int(reqData.GetBool("isUse")),
	})
	if err != nil {
		log.Instance().Errorfln("[Menu IsUse]: %v", err)
	}
	success := err == nil && row > 0
	r.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}
