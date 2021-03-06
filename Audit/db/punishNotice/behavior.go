package db_punishNotice

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addBehavior(tx *gdb.TX, list g.List) (int, error) {
	res, err := tx.BatchInsert(table.PunishNoticeBehavior, list, 5)
	id, _ := res.LastInsertId()
	return int(id), err
}

func updateBehavior(tx *gdb.TX, list g.List) (int, error) {
	res, err := tx.BatchSave(table.PunishNoticeBehavior, list, 5)
	row, _ := res.RowsAffected()
	return int(row), err
}

func EditBehavior(punishNoticeId, basisClauseId int, list [2]g.List) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		row, err = UpdateTX(*tx, punishNoticeId, g.Map{"basis_clause_id": basisClauseId})
		rows += row
	}
	if err == nil {
		row, err = delBehavior(tx, punishNoticeId)
		rows += row
	}
	if err == nil && len(list[0]) > 0 {
		row, err = addBehavior(tx, list[0])
		rows += row
	}
	if err == nil && len(list[1]) > 0 {
		row, err = updateBehavior(tx, list[1])
		rows += row
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		rows = 0
		_ = tx.Rollback()
	}
	return rows, err
}

func delBehavior(tx *gdb.TX, punishNoticeId int) (int, error) {
	r, err := tx.Table(table.PunishNoticeBehavior).Where("punish_notice_id=?", punishNoticeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetBehavior(punishNoticeId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeBehavior + " pn")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.Where("pn.punish_notice_id=?", punishNoticeId)
	sql.And("pn.delete=?", 0)
	sql.OrderBy("pn.id asc")
	sql.Fields("pn.*,u.user_name")
	res, err := sql.All()
	return res.ToList(), err
}

// 统计一个整改通知违规人员的违规数量
func GetBehaviorTotalUserByDraft(draftId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeBehavior + " pn")
	sql.LeftJoin(table.PunishNotice+" p", "pn.id=pn.punish_notice_id")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.Where("p.draft_id=?", draftId)
	sql.Fields("pn.user_id,u.user_name,COUNT(1) AS `sum`")
	r, err := sql.All()
	return r.ToList(), err
}

func CountTopBehavior() (g.List, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeBehavior)
	sql.Fields("content,COUNT(`content`) as sum")
	sql.GroupBy("`content`")
	sql.OrderBy("sum desc")
	sql.Limit(0, 5)
	r, err := sql.All()
	return r.ToList(), err
}

func CountTopDepartment() (g.List, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeBehavior + " pn")
	sql.LeftJoin(table.User+" u", "pn.user_id=u.user_id")
	sql.LeftJoin(table.Department+" d", "u.department_id=d.id")
	sql.Fields("d.name,d.id,COUNT(d.id) as sum")
	sql.GroupBy("d.id")
	sql.OrderBy("sum desc")
	sql.Limit(0, 5)
	r, err := sql.All()
	return r.ToList(), err
}
