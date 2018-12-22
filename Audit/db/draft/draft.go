package db_draft

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func add(tx *gdb.TX, data g.Map) (int, error) {
	res, err := tx.Table(config.DraftTbName).Data(data).Insert()
	id, _ := res.LastInsertId()
	return int(id), err
}

func edit(tx *gdb.TX, id int, data g.Map, where ...g.Map) (int64, error) {
	sql := tx.Table(config.DraftTbName).Data(data).Where("id=?", id)
	sql.And("`delete`=?", 0)
	if len(where) > 0 {
		for k, v := range where[0] {
			sql.And(k, v)
		}
	}
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return row, err
}

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.DraftTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DraftTbName + " d")
	sql.LeftJoin(config.ProgrammeTbName+" p", "d.department_id=p.id")
	sql.LeftJoin(config.ProgrammeTbName+" pq", "d.query_department_id=pq.id")
	sql.Fields("d.*,p.name as department_name,pq.name as query_department_name")
	sql.Where("d.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("d.id desc").Select()
	return r.ToList(), err
}

func Add(draft g.Map, content []g.Map) (int, error) {
	id := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		id, err = add(tx, draft)
	}
	if err == nil && id != 0 {
		_, err = addContent(tx, id, content)
	}
	if err == nil {
		err = tx.Commit()
	} else {
		id = 0
		err = tx.Rollback()
	}
	return id, err
}

func Edit(id int, draft g.Map, content [2][]g.Map, where ...g.Map) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		var r int64 = 0
		r, err = edit(tx, id, draft, where[0])
		rows += int(r)
	}
	if err == nil && rows > 0 {
		_, _ = delContent(tx, id)
		row, err = addContent(tx, id, content[0])
		rows += row
		row, err = updateContent(tx, id, content[1])
		rows += row
	}
	if err == nil {
		err = tx.Commit()
	} else {
		id = 0
		err = tx.Rollback()
	}
	return id, err
}

func Get(id int) (entity.DraftItem, error) {
	db := g.DB()
	draft := entity.DraftItem{}
	fields := []string{
		"d.*",
		"p.name as programme_name",
		"qd.name as query_department_name",
		"dm.name as department_name",
	}
	sql := db.Table(config.DraftTbName + " d")
	sql.LeftJoin(config.ProgrammeTbName+" p", "d.programme_id=p.id")
	sql.LeftJoin(config.DepartmentTbName+" qd", "d.query_department_id=qd.id")
	sql.LeftJoin(config.DepartmentTbName+" dm", "d.department_id=dm.id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("d.delete=?", 0)
	sql.And("d.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&draft)
	}
	return draft, err
}

func Del(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	tx, err := db.Begin()
	if err == nil {
		r, _ := tx.Table(config.DraftTbName).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
		rows, _ = r.RowsAffected()
	}
	if err == nil && rows > 0 {
		_, _ = delContent(tx, id)
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return int(rows), err
}
