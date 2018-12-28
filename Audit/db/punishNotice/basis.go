package db_punishNotice

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addBasis(tx *gdb.TX, punishNoticeId int, programmeBasisIds string) (int, error) {
	programmeBasisIdsArr := strings.Split(programmeBasisIds, ",")
	list := []g.Map{}
	for _, v := range programmeBasisIdsArr {
		programmeBasisId := gconv.Int(v)
		if programmeBasisId != 0 {
			list = append(list, g.Map{"punish_notice_id": punishNoticeId, "programme_basis_id": programmeBasisId})
		}
	}
	if len(list) == 0 {
		return 0, nil
	}
	res, err := tx.BatchInsert(table.PunishNoticeBasis, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func editBasis(punishNoticeId int, programmeBasisIds string) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		row, err = delBasis(tx, punishNoticeId)
		rows += row
	}
	if err == nil {
		row, err = addBasis(tx, punishNoticeId, programmeBasisIds)
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

func updateBasis(tx *gdb.TX, punishNoticeId int, programmeBasisIds string) (int, error) {
	programmeBasisIdsArr := strings.Split(programmeBasisIds, ",")
	programmeBasisId := g.Slice{}
	for _, v := range programmeBasisIdsArr {
		pId := gconv.Int(v)
		if pId != 0 {
			programmeBasisId = append(programmeBasisId, pId)
		}
	}
	if len(programmeBasisId) == 0 {
		return 0, nil
	}
	sql := tx.Table(table.PunishNoticeBasis).Data(g.Map{"delete": 0})
	sql.Where("punish_notice_id=?", punishNoticeId)
	sql.And("programme_basis_id IN ?", programmeBasisId)
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return int(row), err
}

func delBasis(tx *gdb.TX, punishNoticeId int) (int, error) {
	r, err := tx.Table(table.PunishNoticeBasis).Where("punish_notice_id=?", punishNoticeId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetBasis(punishNoticeId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.PunishNoticeBasis + " p")
	sql.LeftJoin(table.ProgrammeBasis+" pb", "p.programme_basis_id=u.id")
	sql.Fields("pb.*,p.*")
	sql.Where("p.punish_notice_id=?", punishNoticeId)
	sql.And("p.delete=?", 0)
	sql.OrderBy("p.id asc")
	res, err := sql.All()
	return res.ToList(), err
}
