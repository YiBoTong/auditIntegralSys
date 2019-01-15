package db_punishNotice

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addScore(tx *gdb.TX, punishNoticeId, cognizanceUserId, score, money int) (int, error) {
	res, err := tx.Table(table.PunishNoticeScore).Data(g.Map{
		"score":              score,
		"money":              money,
		"punish_notice_id":   punishNoticeId,
		"cognizance_user_id": cognizanceUserId,
		"update_time":        util.GetLocalNowTimeStr(),
	}).Insert()
	id, _ := res.LastInsertId()
	return int(id), err
}

func EditScore(punishNoticeId, todoUserId, score, money int) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		row, err = delScore(tx, punishNoticeId)
		rows += row
	}
	if err == nil {
		row, err = addScore(tx, punishNoticeId, todoUserId, score, money)
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

func delScore(tx *gdb.TX, punishNoticeId int) (int, error) {
	r, err := tx.Table(table.PunishNoticeScore).Where("punish_notice_id=?", punishNoticeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetScore(punishNoticeId int) (entity.PunishNoticeScore, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeScore)
	sql.Where("punish_notice_id=?", punishNoticeId)
	sql.And("`delete`=?", 0)
	punishNoticeScore := entity.PunishNoticeScore{}
	res, err := sql.One()
	_ = res.ToStruct(&punishNoticeScore)
	return punishNoticeScore, err
}

func getScoreWidthPunishNotice(punishNoticeId int) (entity.PunishNoticeWidthScore, error) {
	db := g.DB()
	sql := db.Table(table.PunishNotice + " p")
	sql.LeftJoin(table.PunishNoticeScore+" ps", "p.id=ps.punish_notice_id")
	sql.Where("ps.delete=? AND p.id=?", 0, punishNoticeId)
	sql.Fields("p.*,ps.cognizance_user_id,ps.score,ps.update_time")
	punishNoticeWidthScore := entity.PunishNoticeWidthScore{}
	res, err := sql.One()
	_ = res.ToStruct(&punishNoticeWidthScore)
	return punishNoticeWidthScore, err
}
