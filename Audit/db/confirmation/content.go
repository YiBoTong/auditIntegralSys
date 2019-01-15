package db_confirmation

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addContent(tx *gdb.TX, confirmationId int, list []g.Map) (int, error) {
	l := len(list)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		list[i]["confirmation_id"] = confirmationId
	}
	res, err := tx.BatchInsert(table.ConfirmationContent, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateContent(tx *gdb.TX, confirmationId int, list []g.Map) (int, error) {
	l := len(list)
	if l == 0 {
		return 0, nil
	}
	for i := 0; i < l; i++ {
		list[i]["confirmation_id"] = confirmationId
		list[i]["delete"] = 0
	}
	res, err := tx.BatchSave(table.ConfirmationContent, list, 5)
	row, _ := res.RowsAffected()
	return int(row), err
}

func delContent(tx *gdb.TX, confirmationId int) (int, error) {
	r, err := tx.Table(table.ConfirmationContent).Where("confirmation_id=?", confirmationId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetContent(confirmationId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ConfirmationContent).Where("confirmation_id=?", confirmationId)
	sql.And("`delete`=?", 0)
	sql.OrderBy("`order` asc")
	res, err := sql.All()
	return res.ToList(), err
}

// 获取违规分类最多的事实确认书
func GetTopContentTitle(num int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.ConfirmationContent)
	sql.Where("`type`=? AND `delete`=?", "title", 0)
	sql.GroupBy("confirmation_id")
	sql.OrderBy("sum")
	sql.Fields("confirmation_id,COUNT(1) AS `sum`")
	sql.Limit(0, num)
	res, err := sql.Select()
	return res.ToList(), err
}
