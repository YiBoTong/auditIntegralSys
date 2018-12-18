package handler

import (
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"fmt"
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

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"title": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, "p."+k+":"+k, gconv.String(v))
	}

	count, err := db_programme.Count(searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_programme.List(offset, size, listSearchMap)
		for _, v := range listData {
			programmeItem := entity.ProgrammeItem{}
			err = gconv.Struct(v, &programmeItem)
			if err == nil {
				rspData = append(rspData, programmeItem)
			} else {
				break
			}
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
	reqBasis := reqData.GetJson("basis")
	reqContent := reqData.GetJson("content")
	reqStep := reqData.GetJson("step")
	reqBusiness := reqData.GetJson("business")
	reqEmphases := reqData.GetJson("emphases")
	reqUserList := reqData.GetJson("userList")

	addProgramme := map[string]interface{}{
		"title":               "string",
		"key":                 "string",
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
		"update_time":         "nowTime", // 当前时间
	}
	addBasis := map[string]interface{}{
		"clause_id": "int",
		"order":     "int",
		"content":   "string",
	}
	addContent := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addStep := map[string]interface{}{
		"order":   "int",
		"content": "string",
		"type":    "string",
	}
	addBusiness := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addEmphases := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addUserList := map[string]interface{}{
		"user_id": "int",
		"job":     "string",
		"title":   "string",
		"task":    "string",
		"order":   "int",
	}

	programme := g.Map{}
	basis := []g.Map{}
	content := []g.Map{}
	step := []g.Map{}
	business := []g.Map{}
	emphases := []g.Map{}
	userList := []g.Map{}

	for k, v := range addProgramme {
		util.GetSqlMapByReqJson(programme, *reqData, k, gconv.String(v))
	}

	if reqBasisLen := len(reqBasis.ToArray()); reqBasisLen > 0 {
		for basisIndex := 0; basisIndex < reqBasisLen; basisIndex++ {
			item := g.Map{}
			reqBasisItem := reqBasis.GetJson(gconv.String(basisIndex))
			for k, v := range addBasis {
				util.GetSqlMapByReqJson(item, *reqBasisItem, k, gconv.String(v))
			}
			basis = append(basis, item)
		}
	}

	if reqContentLen := len(reqContent.ToArray()); reqContentLen > 0 {
		for contentIndex := 0; contentIndex < reqContentLen; contentIndex++ {
			item := g.Map{}
			reqContentItem := reqContent.GetJson(gconv.String(contentIndex))
			for k, v := range addContent {
				util.GetSqlMapByReqJson(item, *reqContentItem, k, gconv.String(v))
			}
			content = append(content, item)
		}
	}

	if reqStepLen := len(reqStep.ToArray()); reqStepLen > 0 {
		for stepIndex := 0; stepIndex < reqStepLen; stepIndex++ {
			item := g.Map{}
			reqStepItem := reqStep.GetJson(gconv.String(stepIndex))
			for k, v := range addStep {
				util.GetSqlMapByReqJson(item, *reqStepItem, k, gconv.String(v))
			}
			step = append(step, item)
		}
	}

	if reqBusinessLen := len(reqBusiness.ToArray()); reqBusinessLen > 0 {
		for businessIndex := 0; businessIndex < reqBusinessLen; businessIndex++ {
			item := g.Map{}
			reqBusinessItem := reqBusiness.GetJson(gconv.String(businessIndex))
			for k, v := range addBusiness {
				util.GetSqlMapByReqJson(item, *reqBusinessItem, k, gconv.String(v))
			}
			business = append(business, item)
		}
	}

	if reqEmphasesLen := len(reqEmphases.ToArray()); reqEmphasesLen > 0 {
		for emphasesIndex := 0; emphasesIndex < reqEmphasesLen; emphasesIndex++ {
			item := g.Map{}
			reqEmphasesItem := reqEmphases.GetJson(gconv.String(emphasesIndex))
			for k, v := range addEmphases {
				util.GetSqlMapByReqJson(item, *reqEmphasesItem, k, gconv.String(v))
			}
			emphases = append(emphases, item)
		}
	}

	if reqUserListLen := len(reqEmphases.ToArray()); reqUserListLen > 0 {
		for userIndex := 0; userIndex < reqUserListLen; userIndex++ {
			item := g.Map{}
			reqUserItem := reqUserList.GetJson(gconv.String(userIndex))
			for k, v := range addUserList {
				util.GetSqlMapByReqJson(item, *reqUserItem, k, gconv.String(v))
			}
			userList = append(userList, item)
		}
	}

	fmt.Println(programme)
	fmt.Println(basis)
	fmt.Println(content)
	fmt.Println(step)
	fmt.Println(business)
	fmt.Println(emphases)
	fmt.Println(userList)
	id, err := db_programme.Add(programme, basis, content, step, business, emphases, userList)
	if err != nil {
		log.Instance().Errorfln("[Programme Add]: %v", err)
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

func (c *Programme) Get() {
	id := c.Request.GetQueryInt("id")
	basis := []entity.ProgrammeBasis{}
	content := []entity.ProgrammeContent{}
	step := []entity.ProgrammeStep{}
	business := []entity.ProgrammeBusiness{}
	emphases := []entity.ProgrammeEmphases{}
	userList := []entity.ProgrammeUser{}

	programme, err := db_programme.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Programme Get]: %v", err)
	}

	if programme.Id != 0 {
		basisList, _ := db_programme.GetBasis(id)
		contentList, _ := db_programme.GetContent(id)
		stepList, _ := db_programme.GetStep(id)
		businessList, _ := db_programme.GetBusiness(id)
		emphasesList, _ := db_programme.GetEmphases(id)
		userListList, _ := db_programme.GetUser(id)

		for _, bv := range basisList {
			item := entity.ProgrammeBasis{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				basis = append(basis, item)
			}
		}

		for _, cv := range contentList {
			item := entity.ProgrammeContent{}
			if ok := gconv.Struct(cv, &item); ok == nil {
				content = append(content, item)
			}
		}

		for _, sv := range stepList {
			item := entity.ProgrammeStep{}
			if ok := gconv.Struct(sv, &item); ok == nil {
				step = append(step, item)
			}
		}

		for _, bsv := range businessList {
			item := entity.ProgrammeBusiness{}
			if ok := gconv.Struct(bsv, &item); ok == nil {
				business = append(business, item)
			}
		}

		for _, ev := range emphasesList {
			item := entity.ProgrammeEmphases{}
			if ok := gconv.Struct(ev, &item); ok == nil {
				emphases = append(emphases, item)
			}
		}

		for _, uv := range userListList {
			item := entity.ProgrammeUser{}
			if ok := gconv.Struct(uv, &item); ok == nil {
				userList = append(userList, item)
			}
		}
	}

	success := err == nil && programme.Id != 0
	c.Response.WriteJson(app.Response{
		Data: entity.Programme{
			ProgrammeItem: programme,
			Basis:         basis,
			Content:       content,
			Step:          step,
			Business:      business,
			Emphases:      emphases,
			UserList:      userList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

func (c *Programme) Edit() {
	reqData := c.Request.GetJson()
	id := reqData.GetInt("id")

	reqBasis := reqData.GetJson("basis")
	reqContent := reqData.GetJson("content")
	reqStep := reqData.GetJson("step")
	reqBusiness := reqData.GetJson("business")
	reqEmphases := reqData.GetJson("emphases")
	reqUserList := reqData.GetJson("userList")

	addProgramme := map[string]interface{}{
		"title":               "string",
		"key":                 "string",
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
		"update_time":         "nowTime", // 当前时间
	}
	addBasis := map[string]interface{}{
		"clause_id": "int",
		"order":     "int",
		"content":   "string",
	}
	addContent := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addStep := map[string]interface{}{
		"order":   "int",
		"content": "string",
		"type":    "string",
	}
	addBusiness := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addEmphases := map[string]interface{}{
		"order":   "int",
		"content": "string",
	}
	addUserList := map[string]interface{}{
		"user_id": "int",
		"job":     "string",
		"title":   "string",
		"task":    "string",
		"order":   "int",
	}

	programme := g.Map{}
	basis := [][]g.Map{}
	content := [][]g.Map{}
	step := [][]g.Map{}
	business := [][]g.Map{}
	emphases := [][]g.Map{}
	userList := [][]g.Map{}

	for k, v := range addProgramme {
		util.GetSqlMapByReqJson(programme, *reqData, k, gconv.String(v))
	}

	if reqBasisLen := len(reqBasis.ToArray()); reqBasisLen > 0 {
		for basisIndex := 0; basisIndex < reqBasisLen; basisIndex++ {
			item := g.Map{}
			reqBasisItem := reqBasis.GetJson(gconv.String(basisIndex))
			for k, v := range addBasis {
				util.GetSqlMapByReqJson(item, *reqBasisItem, k, gconv.String(v))
			}
			basis[1] = append(basis[1], item)
		}
	}

	if reqContentLen := len(reqContent.ToArray()); reqContentLen > 0 {
		for contentIndex := 0; contentIndex < reqContentLen; contentIndex++ {
			item := g.Map{}
			reqContentItem := reqContent.GetJson(gconv.String(contentIndex))
			for k, v := range addContent {
				util.GetSqlMapByReqJson(item, *reqContentItem, k, gconv.String(v))
			}
			content[1] = append(content[1], item)
		}
	}

	if reqStepLen := len(reqStep.ToArray()); reqStepLen > 0 {
		for stepIndex := 0; stepIndex < reqStepLen; stepIndex++ {
			item := g.Map{}
			reqStepItem := reqStep.GetJson(gconv.String(stepIndex))
			for k, v := range addStep {
				util.GetSqlMapByReqJson(item, *reqStepItem, k, gconv.String(v))
			}
			step[1] = append(step[1], item)
		}
	}

	if reqBusinessLen := len(reqBusiness.ToArray()); reqBusinessLen > 0 {
		for businessIndex := 0; businessIndex < reqBusinessLen; businessIndex++ {
			item := g.Map{}
			reqBusinessItem := reqBusiness.GetJson(gconv.String(businessIndex))
			for k, v := range addBusiness {
				util.GetSqlMapByReqJson(item, *reqBusinessItem, k, gconv.String(v))
			}
			business[1] = append(business[1], item)
		}
	}

	if reqEmphasesLen := len(reqEmphases.ToArray()); reqEmphasesLen > 0 {
		for emphasesIndex := 0; emphasesIndex < reqEmphasesLen; emphasesIndex++ {
			item := g.Map{}
			reqEmphasesItem := reqEmphases.GetJson(gconv.String(emphasesIndex))
			for k, v := range addEmphases {
				util.GetSqlMapByReqJson(item, *reqEmphasesItem, k, gconv.String(v))
			}
			emphases[1] = append(emphases[1], item)
		}
	}

	if reqUserListLen := len(reqEmphases.ToArray()); reqUserListLen > 0 {
		for userIndex := 0; userIndex < reqUserListLen; userIndex++ {
			item := g.Map{}
			reqUserItem := reqUserList.GetJson(gconv.String(userIndex))
			for k, v := range addUserList {
				util.GetSqlMapByReqJson(item, *reqUserItem, k, gconv.String(v))
			}
			userList[1] = append(userList[1], item)
		}
	}

	fmt.Println(programme)
	fmt.Println(basis)
	fmt.Println(content)
	fmt.Println(step)
	fmt.Println(business)
	fmt.Println(emphases)
	fmt.Println(userList)

	rows, err := db_programme.Edit(id, programme, basis, content, step, business, emphases, userList)

	if err != nil {
		log.Instance().Errorfln("[Programme Edit]: %v", err)
	}
	success := err == nil && rows > 0
	c.Response.WriteJson(app.Response{
		Data: id,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.EditStr, !success),
		},
	})
}

func (c *Programme) Delete() {
	typeId := c.Request.GetQueryInt("id")
	rows, err := db_programme.Del(typeId)
	if err != nil {
		log.Instance().Errorfln("[Programme Delete]: %v", err)
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
