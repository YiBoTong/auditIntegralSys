package handler

import (
	"auditIntegralSys/SystemSetup/db/rbac"
	"auditIntegralSys/Worker/fun"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Rbac struct {
	gmvc.Controller
}

func (r *Rbac) Edit() {
	reqData := r.Request.GetJson()
	key := reqData.GetString("key")
	rbac := reqData.GetJson("rbac")
	rbacLen := len(rbac.ToArray())
	msg := ""
	rows := 0
	rbcas :=[]g.Map{}
	err := error(nil)
	if rbacLen > 0 {
		for i := 0; i < rbacLen; i++ {
			rbcas = append(rbcas, g.Map{
				"key":      key,
				"menu_id":  rbac.GetInt(gconv.String(i) + ".menuId"),
				"is_read":  gconv.Int(rbac.GetBool(gconv.String(i) + ".isRead")),
				"is_write": gconv.Int(rbac.GetBool(gconv.String(i) + ".isWrite")),
			})
		}
		_, err = db_rbac.Del(key)
		if err == nil {
			rows, err = db_rbac.Add(rbcas)
		}
	} else {
		msg = config.RbacStr + config.MastHasOneStr
		err = errors.New(msg)
	}
	if err != nil {
		log.Instance().Errorfln("[Rbac Edit]: %v", err)
	}
	success := err == nil && rows > 0
	if msg == "" {
		msg = config.GetTodoResMsg(config.RbacStr+config.EditStr, !success)
	}
	r.Response.WriteJson(app.Response{
		Data: rows,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   msg,
		},
	})
}

func (c *Rbac) Get() {
	key := c.Request.GetQueryString("key")
	rbacMenu, err := fun.GetAllMenuRbac(-1, key)
	if err != nil {
		log.Instance().Errorfln("[Rbac Get]: %v", err)
	}
	c.Response.WriteJson(app.Response{
		Data: rbacMenu,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.GetStr, err != nil),
		},
	})
}

func (c *Rbac) Delete() {
	db_rbac.RemoveOldData()
	c.Response.WriteJson(app.Response{
		Data: "清除废弃数据",
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   "",
		},
	})
}
