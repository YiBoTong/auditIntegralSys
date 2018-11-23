package handler

import (
	"auditIntegralSys/SystemSetup/db/dictionaries"
	"auditIntegralSys/SystemSetup/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
)

type Dictionaries struct {
	gmvc.Controller
}

func (c *Dictionaries) List() {
	reqData := c.Request.GetJson()
	var rspData []entity.DictionaryType
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	title := search.GetString("title")
	key := search.GetString("key")
	userId := search.GetInt("userId")

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	if title != "" {
		searchMap["title"] = title
		listSearchMap["d.title"] = title
	}

	if key != "" {
		searchMap["'key'"] = key
		listSearchMap["d.key"] = key
	}

	if userId != 0 {
		searchMap["user_id"] = userId
		listSearchMap["d.user_id"] = userId
	}

	count, err := db_dictionaries.GetDictionaryTypeCount(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_dictionaries.GetDictionaryTypes(offset, size, listSearchMap)
		for _, v := range listData {
			rspData = append(rspData, entity.DictionaryType{
				Id:         gconv.Int(v["id"]),
				TypeId:     gconv.Int(v["type_id"]),
				Key:        gconv.String(v["key"]),
				Title:      gconv.String(v["title"]),
				IsUse:      gconv.Bool(v["is_use"]),
				UpdateTime: gconv.String(v["update_time"]),
				UserId:     gconv.Int(v["user_id"]),
				UserName:   gconv.String(v["user_name"]),
				Describe:   gconv.String(v["describe"]),
			})
		}
	}
	if err != nil {
		log.Instance().Errorf("[Dictionaries List]: %v", err)
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

func (c *Dictionaries) Add() {
	reqData := c.Request.GetJson()
	reqDictionaries := reqData.GetJson("dictionaries")

	var dictionaries []g.Map
	dictionaryType := g.Map{
		"type_id":     reqData.GetInt("typeId"),
		"key":         reqData.GetString("key"),
		"title":       reqData.GetString("title"),
		"is_use":      gconv.Int(reqData.GetBool("isUse")),
		"user_id":     reqData.GetInt("userId"),
		"update_time": util.GetLocalNowTimeStr(),
		"describe":    reqData.GetString("describe"),
	}

	id, err := db_dictionaries.AddDictionaryType(dictionaryType)
	dictionaryLen := len(reqDictionaries.ToArray())
	if err == nil && dictionaryLen > 0 {
		for i := 0; i < dictionaryLen; i++ {
			dictionaries = append(dictionaries, g.Map{
				"type_id":  id,
				"key":      reqDictionaries.GetString(gconv.String(i) + ".key"),
				"value":    reqDictionaries.GetString(gconv.String(i) + ".value"),
				"order":    reqDictionaries.GetInt(gconv.String(i) + ".order"),
				"describe": reqDictionaries.GetString(gconv.String(i) + ".describe"),
			})
		}
		_, err = db_dictionaries.AddDictionaries(dictionaries)
		if err != nil {
			_, _ = db_dictionaries.DelDictionaryType(id)
		}
	}
	if err != nil {
		log.Instance().Errorf("[Dictionaries Add]: %v", err)
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

func (c *Dictionaries) Get() {
	typeId := c.Request.GetQueryInt("id")
	var dictionaries = []entity.Dictionary{}
	dictionaryType, err := db_dictionaries.GetDictionaryType(typeId)
	if err == nil {
		var listData []map[string]interface{}
		listData, err = db_dictionaries.GetDictionaries(typeId)
		for _, v := range listData {
			dictionaries = append(dictionaries, entity.Dictionary{
				Id:       gconv.Int(v["id"]),
				TypeId:   gconv.Int(v["type_id"]),
				Key:      gconv.String(v["key"]),
				Value:    gconv.String(v["value"]),
				Order:    gconv.Int(v["order"]),
				Describe: gconv.String(v["describe"]),
			})
		}
	}
	if err != nil {
		log.Instance().Errorf("[Dictionaries Get]: %v", err)
	}
	success := err == nil && dictionaryType.Id > 0
	c.Response.WriteJson(app.Response{
		Data: entity.DictionaryTypeRes{
			DictionaryType: dictionaryType,
			Dictionaries:   dictionaries,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (c *Dictionaries) Edit() {
	reqData := c.Request.GetJson()
	typeId := reqData.GetInt("id")
	reqDictionaries := reqData.GetJson("dictionaries")

	var addDictionaries []g.Map
	var updateDictionaries []g.Map
	var updateDictionaryIds []int
	dictionaryType := g.Map{
		"type_id":     reqData.GetInt("typeId"),
		"key":         reqData.GetString("key"),
		"title":       reqData.GetString("title"),
		"is_use":      gconv.Int(reqData.GetBool("isUse")),
		"user_id":     reqData.GetInt("userId"),
		"update_time": util.GetLocalNowTimeStr(),
		"describe":    reqData.GetString("describe"),
	}

	rows, err := db_dictionaries.UpdateDictionaryType(typeId, dictionaryType)
	dictionaryLen := len(reqDictionaries.ToArray())
	if err == nil && rows > 0 && dictionaryLen > 0 {
		for i := 0; i < dictionaryLen; i++ {
			id := reqDictionaries.GetInt(gconv.String(i) + ".id")
			if id > 1 {
				updateDictionaries = append(updateDictionaries, g.Map{
					"id":       id,
					"type_id":  typeId,
					"key":      reqDictionaries.GetString(gconv.String(i) + ".key"),
					"value":    reqDictionaries.GetString(gconv.String(i) + ".value"),
					"order":    reqDictionaries.GetInt(gconv.String(i) + ".order"),
					"describe": reqDictionaries.GetString(gconv.String(i) + ".describe"),
				})
				updateDictionaryIds = append(updateDictionaryIds, id)
			} else {
				addDictionaries = append(addDictionaries, g.Map{
					"type_id":  typeId,
					"key":      reqDictionaries.GetString(gconv.String(i) + ".key"),
					"value":    reqDictionaries.GetString(gconv.String(i) + ".value"),
					"order":    reqDictionaries.GetInt(gconv.String(i) + ".order"),
					"describe": reqDictionaries.GetString(gconv.String(i) + ".describe"),
				})
			}
		}
		_, err = db_dictionaries.UpdateDictionaries(typeId, addDictionaries, updateDictionaries, updateDictionaryIds)
	}
	if err != nil {
		log.Instance().Errorf("[Dictionaries Edit]: %v", err)
	}
	success := err == nil && rows > 0
	c.Response.WriteJson(app.Response{
		Data: typeId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}

func (c *Dictionaries) IsUse() {
	typeId := c.Request.GetQueryInt("id")
	isUse := c.Request.GetQueryBool("isUse")
	rows, err := db_dictionaries.UpdateDictionaryType(typeId, g.Map{"is_use": gconv.Int(isUse)})
	if err != nil {
		log.Instance().Errorf("[Dictionaries IsUse]: %v", err)
	}
	success := err == nil && rows > 0
	c.Response.WriteJson(app.Response{
		Data: typeId,
		Status: app.Status{
			Code:  0,
			Error: success,
			Msg:   config.GetTodoResMsg(config.ChangeState, success),
		},
	})
}

func (c *Dictionaries) Delete() {
	typeId := c.Request.GetQueryInt("id")
	rows, err := db_dictionaries.DelDictionaryType(typeId)
	if err != nil {
		log.Instance().Errorf("[Dictionaries Delete]: %v", err)
	}
	success := err == nil && rows > 0
	c.Response.WriteJson(app.Response{
		Data: typeId,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.DelStr, !success),
		},
	})
}
