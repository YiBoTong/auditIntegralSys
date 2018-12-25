package db_draft

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addContent(tx *gdb.TX, draftId int, list []g.Map) (int, error) {
	l := len(list)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		list[i]["draft_id"] = draftId
	}
	res, err := tx.BatchInsert(config.DraftContentTbName, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateContent(tx *gdb.TX, draftId int, list []g.Map) (int, error) {
	l := len(list)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		list[i]["draft_id"] = draftId
		list[i]["delete"] = 0
	}
	res, err := tx.BatchSave(config.DraftContentTbName, list, 5)
	row, _ := res.RowsAffected()
	return int(row), err
}

func delContent(tx *gdb.TX, draftId int) (int, error) {
	r, err := tx.Table(config.DraftContentTbName).Where("draft_id=?", draftId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetContent(draftId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DraftContentTbName).Where("draft_id=?", draftId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}
