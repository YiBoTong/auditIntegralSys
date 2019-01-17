package handler

import (
	"auditIntegralSys/Audit/db/auditReport"
	"auditIntegralSys/Audit/db/confirmation"
	"auditIntegralSys/Audit/db/integral"
	"auditIntegralSys/Audit/db/programme"
	"auditIntegralSys/Audit/db/punishNotice"
	"auditIntegralSys/Audit/db/statistical"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	"auditIntegralSys/Org/db/user"
	"auditIntegralSys/_public/app"
	"auditIntegralSys/_public/config"
	"auditIntegralSys/_public/log"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"time"
)

type Statistical struct {
	gmvc.Controller
}

func (r *Statistical) List() {
	reqData := r.Request.GetJson()
	rspData := []entity.StatisticalListItem{}
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	thisUserId := util.GetUserIdByRequest(r.Cookie)

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"project_name": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	thisUserInfo, _ := db_user.GetUser(thisUserId)
	count, err := db_auditReport.Count(thisUserInfo, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_auditReport.List(thisUserInfo, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.StatisticalListItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Draft List]: %v", err)
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

func (r *Statistical) Detailed() {
	reqData := r.Request.GetJson()
	rspData := []entity.StatisticalDetailedItem{}
	// 分页
	pager := reqData.GetJson("page")
	page := pager.GetInt("page")
	size := pager.GetInt("size")
	offset := (page - 1) * size
	search := reqData.GetJson("search")
	thisUserId := util.GetUserIdByRequest(r.Cookie)

	searchMap := g.Map{}
	listSearchMap := g.Map{}

	searchItem := map[string]interface{}{
		"user_name": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
		// p.title:title String
		util.GetSearchMapByReqJson(listSearchMap, *search, k, gconv.String(v))
	}

	thisUserInfo, _ := db_user.GetUser(thisUserId)
	count, err := db_statistical.DetailedCount(thisUserInfo, searchMap)
	if err == nil && offset <= count {
		var listData []map[string]interface{}
		listData, err = db_statistical.DetailedList(thisUserInfo, offset, size, listSearchMap)
		for _, v := range listData {
			item := entity.StatisticalDetailedItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				rspData = append(rspData, item)
			}
		}
	}
	if err != nil {
		log.Instance().Errorfln("[Draft List]: %v", err)
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

func (r *Statistical) Detailed_total() {
	reqData := r.Request.GetJson()
	// 分页
	search := reqData.GetJson("search")
	thisUserId := util.GetUserIdByRequest(r.Cookie)

	searchMap := g.Map{}

	searchItem := map[string]interface{}{
		"user_name": "string",
	}

	for k, v := range searchItem {
		// title String
		util.GetSearchMapByReqJson(searchMap, *search, k, gconv.String(v))
	}

	SumMoney := 0
	SumScore := 0
	thisUserInfo, _ := db_user.GetUser(thisUserId)
	sumList, err := db_statistical.DetailedSum(thisUserInfo, searchMap)
	for _, v := range sumList {
		item := entity.StatisticalDetailedTotal{}
		if ok := gconv.Struct(v, &item); ok == nil {
			SumScore += item.SumScore
			SumMoney += item.SumMoney
		}
	}

	if err != nil {
		log.Instance().Errorfln("[Draft List]: %v", err)
	}
	r.Response.WriteJson(app.Response{
		Data: entity.StatisticalDetailedTotal{
			SumScore: SumScore,
			SumMoney: SumMoney,
		},
		Status: app.Status{
			Code:  0,
			Error: err != nil,
			Msg:   config.GetTodoResMsg(config.GetStr, err != nil),
		},
	})
}

func (r *Statistical) Get() {
	id := r.Request.GetQueryInt("id")
	BusinessList := []entity.ProgrammeBusiness{}

	StatisticalListItem, err := db_statistical.Get(id)

	if err != nil {
		log.Instance().Errorfln("[Rectify Get]: %v", err)
	}

	if StatisticalListItem.Id != 0 {
		basisList, _ := db_programme.GetBusiness(StatisticalListItem.ProgrammeId)
		for _, bv := range basisList {
			item := entity.ProgrammeBusiness{}
			if ok := gconv.Struct(bv, &item); ok == nil {
				BusinessList = append(BusinessList, item)
			}
		}
	}

	success := err == nil && StatisticalListItem.Id != 0
	r.Response.WriteJson(app.Response{
		Data: entity.Statistical{
			StatisticalListItem: StatisticalListItem,
			BusinessList:        BusinessList,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 年度积分
func (r *Statistical) Get_my_score_by_year() {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	rspData := []entity.StatisticalMyNumByYear{}

	err := error(nil)
	thisYear := time.Now().Year()

	for yearIndex := 0; yearIndex < 2; yearIndex++ {
		year := thisYear - yearIndex
		yearInfo := []entity.StatisticalMyNumByYearItem{}
		for month := 1; month < 13; month++ {
			begin, end := fun.GetBetweenOneMonth(year, month)
			score, _ := db_integral.GetBetweenTimeScore(thisUserId, begin, end)
			item := entity.StatisticalMyNumByYearItem{
				Month: month,
				Num:   score,
			}
			yearInfo = append(yearInfo, item)
		}
		yearItem := entity.StatisticalMyNumByYear{
			Year: year,
			Info: yearInfo,
		}
		rspData = append(rspData, yearItem)
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 年度违规行为
func (r *Statistical) Get_my_behavior_by_year() {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	rspData := []entity.StatisticalMyNumByYear{}

	err := error(nil)
	thisYear := time.Now().Year()

	for yearIndex := 0; yearIndex < 2; yearIndex++ {
		year := thisYear - yearIndex
		yearInfo := []entity.StatisticalMyNumByYearItem{}
		for month := 1; month < 13; month++ {
			begin, end := fun.GetBetweenOneMonth(year, month)
			score, _ := db_confirmation.GetBetweenTimeNum(thisUserId, begin, end)
			item := entity.StatisticalMyNumByYearItem{
				Month: month,
				Num:   score,
			}
			yearInfo = append(yearInfo, item)
		}
		yearItem := entity.StatisticalMyNumByYear{
			Year: year,
			Info: yearInfo,
		}
		rspData = append(rspData, yearItem)
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 年度事实确认书
func (r *Statistical) Get_my_confirmation_by_year() {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	rspData := []entity.StatisticalMyNumByYear{}

	err := error(nil)
	thisYear := time.Now().Year()

	for yearIndex := 0; yearIndex < 2; yearIndex++ {
		year := thisYear - yearIndex
		yearInfo := []entity.StatisticalMyNumByYearItem{}
		for month := 1; month < 13; month++ {
			begin, end := fun.GetBetweenOneMonth(year, month)
			score, _ := db_confirmation.GetBetweenTimeNum(thisUserId, begin, end)
			item := entity.StatisticalMyNumByYearItem{
				Month: month,
				Num:   score,
			}
			yearInfo = append(yearInfo, item)
		}
		yearItem := entity.StatisticalMyNumByYear{
			Year: year,
			Info: yearInfo,
		}
		rspData = append(rspData, yearItem)
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 年度惩罚通知书
func (r *Statistical) Get_my_punish_notice_by_year() {
	thisUserId := util.GetUserIdByRequest(r.Cookie)
	rspData := []entity.StatisticalMyNumByYear{}

	err := error(nil)
	thisYear := time.Now().Year()

	for yearIndex := 0; yearIndex < 2; yearIndex++ {
		year := thisYear - yearIndex
		yearInfo := []entity.StatisticalMyNumByYearItem{}
		for month := 1; month < 13; month++ {
			begin, end := fun.GetBetweenOneMonth(year, month)
			score, _ := db_confirmation.GetBetweenTimeNum(thisUserId, begin, end)
			item := entity.StatisticalMyNumByYearItem{
				Month: month,
				Num:   score,
			}
			yearInfo = append(yearInfo, item)
		}
		yearItem := entity.StatisticalMyNumByYear{
			Year: year,
			Info: yearInfo,
		}
		rspData = append(rspData, yearItem)
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: rspData,
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 违规行为top
func (r *Statistical) Get_top_department() {
	//thisUserId := util.GetUserIdByRequest(r.Cookie)
	Behavior := []entity.StatisticalTopBehaviorContentItem{}
	Department := []entity.StatisticalTopBehaviorDepartmentItem{}
	Score := []entity.StatisticalMyNumByYearItem{}

	behavior, _ := db_punishNotice.CountTopBehavior()
	department, _ := db_punishNotice.CountTopDepartment()

	for _, bv := range behavior {
		item := entity.StatisticalTopBehaviorContentItem{}
		if ok := gconv.Struct(bv, &item); ok == nil {
			Behavior = append(Behavior, item)
		}
	}

	for _, sv := range department {
		item := entity.StatisticalTopBehaviorDepartmentItem{}
		if ok := gconv.Struct(sv, &item); ok == nil {
			Department = append(Department, item)
		}
	}

	thisYear := time.Now().Year()
	for month := 1; month < 13; month++ {
		begin, end := fun.GetBetweenOneMonth(thisYear, month)
		score, _ := db_integral.CountMonthScore(begin, end)
		item := entity.StatisticalMyNumByYearItem{
			Month: month,
			Num:   score,
		}
		Score = append(Score, item)
	}

	success := true
	r.Response.WriteJson(app.Response{
		Data: entity.StatisticalTopBehaviorAndDepartment{
			Behavior:   Behavior,
			Department: Department,
			Score:      Score,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}

// 违规行为top
func (r *Statistical) Get_one_statistical_user() {
	//thisUserId := util.GetUserIdByRequest(r.Cookie)
	draftId := r.Request.GetQueryInt("draftId")
	Behavior := []entity.StatisticalOneUserBehaviorItem{}
	Score := []entity.StatisticalOneUserScoreItem{}

	scoreList := g.List{}
	behaviorList, err := db_punishNotice.GetBehaviorTotalUserByDraft(draftId)
	scoreList, err = db_integral.GetScoreTotalUserByDraft(draftId)

	for _, bv := range behaviorList {
		item := entity.StatisticalOneUserBehaviorItem{}
		if ok := gconv.Struct(bv, &item); ok == nil {
			Behavior = append(Behavior, item)
		}
	}

	for _, sv := range scoreList {
		item := entity.StatisticalOneUserScoreItem{}
		if ok := gconv.Struct(sv, &item); ok == nil {
			Score = append(Score, item)
		}
	}

	success := err == nil
	r.Response.WriteJson(app.Response{
		Data: entity.StatisticalOneUser{
			Score:    Score,
			Behavior: Behavior,
		},
		Status: app.Status{
			Code:  0,
			Error: !success,
			Msg:   config.GetTodoResMsg(config.GetStr, !success),
		},
	})
}
