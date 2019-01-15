package db_punishNotice

import (
	"auditIntegralSys/Audit/db/integral"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.Draft+" d", "pn.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")

	sql.Where("pn.delete=?", 0)
	sql.GroupBy("pn.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(authorInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, authorInfo, where).Count()
	return r, err
}

func List(authorInfo entity2.User, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"pn.*",
		"u.user_name",
		"dt.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
		"p.start_time",
		"p.end_time",
		"p.plan_start_time",
		"p.plan_end_time",
	}
	sql := getListSql(db, authorInfo, where)
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("pn.id desc").Select()
	return r.ToList(), err
}

func Get(id int, where ...g.Map) (entity.PunishNoticeItem, error) {
	db := g.DB()
	fields := []string{
		"d.*",
		"pn.*",
		"u.user_name",
		"dt.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
		"p.start_time",
		"p.end_time",
		"p.plan_start_time",
		"p.plan_end_time",
		"c.title as basis_clause_title",
		"c.number as basis_clause_number",
	}
	confirmation := entity.PunishNoticeItem{}
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.Draft+" d", "pn.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dt", "d.department_id=dt.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.LeftJoin(table.Clause+" c", "pn.basis_clause_id=c.id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("pn.delete=?", 0)
	sql.And("pn.id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&confirmation)
	}
	return confirmation, err
}

func GetUsersByConfirmationId(confirmationId int, where ...g.Map) (g.List, error) {
	db := g.DB()
	fields := []string{
		"pn.*",
		"ps.score,ps.update_time",
		"u.user_name",
	}
	sql := db.Table(table.PunishNotice + " pn")
	sql.LeftJoin(table.PunishNoticeScore+" ps", "ps.punish_notice_id=pn.id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("pn.delete=?", 0)
	sql.And("pn.confirmation_id=?", confirmationId)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.OrderBy("pn.id asc").All()
	return r.ToList(), err
}

func Add(tx gdb.TX, confirmationId, draftId int, userIds []int) (int, error) {
	data := g.List{}
	nowTime := util.GetLocalNowTimeStr()
	for _, v := range userIds {
		data = append(data, g.Map{
			"confirmation_id": confirmationId,
			"draft_id":        draftId,
			"user_id":         v,
			"time":            nowTime,
		})
	}
	if len(data) < 1 {
		return 0, nil
	}
	r, err := tx.BatchInsert(table.PunishNotice, data, 3)
	id, _ := r.LastInsertId()
	return int(id), err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.PunishNotice).Data(data)
	sql.Where("`delete`=?", 0)
	sql.And("id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func UpdateTX(tx gdb.TX, id int, data g.Map, where ...g.Map) (int, error) {
	sql := tx.Table(table.PunishNotice).Data(data)
	sql.Where("`delete`=?", 0)
	sql.And("id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Publish(id int, number string, where g.Map) (int, error) {
	db := g.DB()
	rows := 0
	tx, err := db.Begin()
	punishNoticeWidthScore := entity.PunishNoticeWidthScore{}

	if err == nil {
		sql := tx.Table(table.PunishNotice).Data(g.Map{"number": number, "state": state.Publish})
		sql.Where("id=? AND `delete`=?", id, 0).And(where)
		r, _ := sql.Update()
		row, _ := r.RowsAffected()
		rows = int(row)
	}
	if err == nil && rows != 0 {
		punishNoticeWidthScore, err = getScoreWidthPunishNotice(id)
	}
	if err == nil && punishNoticeWidthScore.Id != 0 {
		data := g.Map{
			"cognizance_user_id": punishNoticeWidthScore.CognizanceUserId, // 认定人ID
			"user_id":            punishNoticeWidthScore.UserId,           // 责任人ID
			"draft_id":           punishNoticeWidthScore.DraftId,          // 工作底稿ID
			"punish_notice_id":   id,                                      // 处罚通知ID
			"score":              punishNoticeWidthScore.Score,            // 分数
			"money":              punishNoticeWidthScore.Money,            // 金额
			"time":               punishNoticeWidthScore.UpdateTime,       // 认定时间
		}
		// 分数记录到档案中（积分表）
		_, err = db_integral.AddScore(*tx, data)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return rows, err
}
