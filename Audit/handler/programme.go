package handler

import (
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Programme struct {
	gmvc.Controller
}

func (c *Programme) List() {
	reqData := c.Request.GetJson()
	var rspData []entity.ProgrammeItem
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	title := search.GetString("title")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	if title != "" {
		searchMap["title"] = title
		listSearchMap["d.title"] = title
	}

	count, err := db_programme.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeItem{}
			err := gconv.Struct(v, programmeItem)
			if err == nil {
				rspData = append(rspData, programmeItem)
			}
			//rspData = append(rspData, entity.ProgrammeItem{
			//	Id:                  gconv.Int(v["id"]),
			//	QueryDepartmentId:   0,
			//	QueryDepartmentName: "",
			//	UserId:              gconv.Int(v["user_id"]),
			//	QueryPointId:        0,
			//	QueryPointName:      "",
			//	Purpose:             "",
			//	Type:                "",
			//	StartTime:           "",
			//	EndTime:             "",
			//	PlanStartTime:       "",
			//	PlanEndTime:         "",
			//	DetUserId:           0,
			//	DetUserName:         "",
			//	DetUserContent:      "",
			//	DetUserTime:         "",
			//	AdminUserId:         "",
			//	AdminUserName:       "",
			//	AdminUserContent:    "",
			//	AdminUserTime:       "",
			//	State:               "",
			//})
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Programme List]: %v", err)
	}
	c.Response.WriteJson(app.ListResponse{
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

func (c *Programme) Add() {
	reqData := c.Request.GetJson()
	//reqBasis := reqData.GetJson("basis")
	//reqContent := reqData.GetJson("content")
	//reqStep := reqData.GetJson("step")
	//reqBusiness := reqData.GetJson("business")
	//reqEmphases := reqData.GetJson("emphases")
	//reqUserList := reqData.GetJson("userList")

	addProgramme := map[string]interface{}{
		"query_department_id": "int",
		"user_id":             "int",
		"query_point_id":      "int",
		"purpose":             "string",
		"type":                "string",
		"start_time":          "string",
		"end_time":            "string",
		"plan_start_time":     "string",
		"plan_end_time":       "string",
		"state":               "string",
	}
	programme := g.Map{
		"update_time": util.GetLocalNowTimeStr(),
	}

	for k, v := range addProgramme {
		programme[k] = gconv.Convert(reqData.Get(util.CamelCase(k)), gconv.String(v))
	}

	id := 0
	var err error = nil
	//id, err := db_programme.Add()
	if err != nil {
		log.Instance().Errorfln("[Dictionaries Add]: %v", err)
	}
	c.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.AddStr, err != nil),
		},
	})
}

//func (c *Programme) Get() {
//	typeId := c.Request.GetQueryInt("id")
//	var dictionaries = []entity.Dictionary{}
//	dictionaryType, err := db_dictionaries.GetDictionaryType(typeId)
//	if err == nil {
//		var listData []map[string]interface{}
//		listData, err = db_dictionaries.GetDictionaries(typeId)
//		for _, v := range listData {
//			dictionaries = append(dictionaries, entity.Dictionary{
//				Id:       gconv.Int(v["id"]),
//				TypeId:   gconv.Int(v["type_id"]),
//				Key:      gconv.String(v["key"]),
//				Value:    gconv.String(v["value"]),
//				Order:    gconv.Int(v["order"]),
//				Describe: gconv.String(v["describe"]),
//			})
//		}
//	}
//	if err != nil {
//		log.Instance().Errorfln("[Dictionaries Get]: %v", err)
//	}
//	success := err == nil && dictionaryType.Id != 0
//	c.Response.WriteJson(app.Response{
//		Data: entity.DictionaryTypeRes{
//			DictionaryType: dictionaryType,
//			Dictionaries:   dictionaries,
//		},
//		Status: app.Status{
//			Code:  0,
//			Error: !success,
//			Msg:   config.GetTodoResMsg(config.GetStr, !success),
//		},
//	})
//}
//
//func (c *Programme) Edit() {
//	reqData := c.Request.GetJson()
//	typeId := reqData.GetInt("id")
//	reqDictionaries := reqData.GetJson("dictionaries")
//
//	var addDictionaries []g.Map
//	var updateDictionaries []g.Map
//	var updateDictionaryIds []int
//	dictionaryType := g.Map{
//		"type_id":     reqData.GetInt("typeId"),
//		"key":         reqData.GetString("key"),
//		"title":       reqData.GetString("title"),
//		"is_use":      gconv.Int(reqData.GetBool("isUse")),
//		"user_id":     reqData.GetInt("userId"),
//		"update_time": util.GetLocalNowTimeStr(),
//		"describe":    reqData.GetString("describe"),
//	}
//
//	rows, err := db_dictionaries.UpdateDictionaryType(typeId, dictionaryType)
//	dictionaryLen := len(reqDictionaries.ToArray())
//	if err == nil && rows > 0 && dictionaryLen > 0 {
//		for i := 0; i < dictionaryLen; i++ {
//			id := reqDictionaries.GetInt(gconv.String(i) + ".id")
//			if id != 0 {
//				updateDictionaries = append(updateDictionaries, g.Map{
//					"id":       id,
//					"type_id":  typeId,
//					"key":      reqDictionaries.GetString(gconv.String(i) + ".key"),
//					"value":    reqDictionaries.GetString(gconv.String(i) + ".value"),
//					"order":    reqDictionaries.GetInt(gconv.String(i) + ".order"),
//					"describe": reqDictionaries.GetString(gconv.String(i) + ".describe"),
//				})
//				updateDictionaryIds = append(updateDictionaryIds, id)
//			} else {
//				addDictionaries = append(addDictionaries, g.Map{
//					"type_id":  typeId,
//					"key":      reqDictionaries.GetString(gconv.String(i) + ".key"),
//					"value":    reqDictionaries.GetString(gconv.String(i) + ".value"),
//					"order":    reqDictionaries.GetInt(gconv.String(i) + ".order"),
//					"describe": reqDictionaries.GetString(gconv.String(i) + ".describe"),
//				})
//			}
//		}
//		_, err = db_dictionaries.UpdateDictionaries(typeId, addDictionaries, updateDictionaries, updateDictionaryIds)
//	}
//	if err != nil {
//		log.Instance().Errorfln("[Dictionaries Edit]: %v", err)
//	}
//	success := err == nil && rows > 0
//	c.Response.WriteJson(app.Response{
//		Data: typeId,
//		Status: app.Status{
//			Code:  0,
//			Error: !success,
//			Msg:   config.GetTodoResMsg(config.EditStr, !success),
//		},
//	})
//}
//
//func (c *Programme) Delete() {
//	typeId := c.Request.GetQueryInt("id")
//	rows, err := db_dictionaries.DelDictionaryType(typeId)
//	if err != nil {
//		log.Instance().Errorfln("[Dictionaries Delete]: %v", err)
//	}
//	success := err == nil && rows > 0
//	c.Response.WriteJson(app.Response{
//		Data: typeId,
//		Status: app.Status{
//			Code:  0,
//			Error: !success,
//			Msg:   config.GetTodoResMsg(config.DelStr, !success),
//		},
//	})
//}
