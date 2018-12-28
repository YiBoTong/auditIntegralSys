package db_menu

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
)

func Menus(parentId int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Menu).Where("`delete`=?", 0)
	sql.And("parent_id=?", parentId)
	if len(where) > 0 {
		sql.And(where)
	}
	res, err := sql.OrderBy("`order` asc").All()
	return res.ToList(), err
}

func Add(add g.Map) (int, error) {
	db := g.DB()
	var id int64 = 0
	res, err := db.Table(table.Menu).Data(add).Insert()
	if err == nil {
		id, _ = res.LastInsertId()
	}
	return int(id), err
}

func Update(id int, data g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.Table(table.Menu).Data(data).Where("id=?", id).Update()
	if err == nil {
		rows, err = r.RowsAffected()
	}
	return int(rows), err
}
