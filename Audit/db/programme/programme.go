package db_programme

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func add(tx *gdb.TX, data g.Map) (int, error) {
	res, err := tx.Table(table.Programme).Data(data).Insert()
	id, _ := res.LastInsertId()
	return int(id), err
}

func edit(tx *gdb.TX, id int, data g.Map, where ...g.Map) (int64, error) {
	sql := tx.Table(table.Programme).Data(data).Where("id=?", id)
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
	sql := db.Table(table.Programme).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.Programme + " p")
	sql.Where("p.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("p.id desc").Select()
	return r.ToList(), err
}

func Add(programme g.Map, basis, content, step, business, emphases, user []g.Map) (int, error) {
	id := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		id, err = add(tx, programme)
	}
	if err == nil && id != 0 {
		_, err = addBasis(tx, id, basis)
	}
	if err == nil && id != 0 {
		_, err = addContent(tx, id, content)
	}
	if err == nil && id != 0 {
		_, err = addStep(tx, id, step)
	}
	if err == nil && id != 0 {
		_, err = addBusiness(tx, id, business)
	}
	if err == nil && id != 0 {
		_, err = addEmphases(tx, id, emphases)
	}
	if err == nil && id != 0 {
		_, err = addUser(tx, id, user)
	}
	if err == nil {
		err = tx.Commit()
	} else {
		id = 0
		err = tx.Rollback()
	}
	return id, err
}

func Edit(id int, programme g.Map, basis, content, step, business, emphases, user [2][]g.Map, where ...g.Map) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		var r int64 = 0
		r, err = edit(tx, id, programme, where[0])
		rows += int(r)
	}
	if err == nil && rows > 0 {
		if err == nil {
			_, _ = delBasis(tx, id)
			row, err = addBasis(tx, id, basis[0])
			rows += row
			row, err = updateBasis(tx, id, basis[1])
			rows += row
		}
		if err == nil {
			_, _ = delContent(tx, id)
			row, err = addContent(tx, id, content[0])
			rows += row
			row, err = updateContent(tx, id, content[1])
			rows += row
		}
		if err == nil {
			_, _ = delStep(tx, id)
			row, err = addStep(tx, id, step[0])
			rows += row
			row, err = updateStep(tx, id, step[1])
			rows += row
		}
		if err == nil {
			_, _ = delBusiness(tx, id)
			row, err = addBusiness(tx, id, business[0])
			rows += row
			row, err = updateBusiness(tx, id, business[1])
			rows += row
		}
		if err == nil {
			_, _ = delEmphases(tx, id)
			row, err = addEmphases(tx, id, emphases[0])
			rows += row
			row, err = updateEmphases(tx, id, emphases[1])
			rows += row
		}
		if err == nil {
			_, _ = delUser(tx, id)
			row, err = addUser(tx, id, user[0])
			rows += row
			row, err = updateUser(tx, id, user[1])
			rows += row
		}
	}
	if err == nil {
		err = tx.Commit()
	} else {
		rows = 0
		err = tx.Rollback()
	}
	return rows, err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Programme).Data(data).Where("`delete`=?", 0).And("id=?", id)
	if len(where) > 0 {
		for k, v := range where[0] {
			sql.And(k, v)
		}
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	fmt.Println(row)
	return int(row), err
}

func Get(id int) (entity.ProgrammeItem, error) {
	db := g.DB()
	programme := entity.ProgrammeItem{}
	fields := []string{
		"p.*",
		"qd.name as query_department_name",
		"qp.name as query_point_name",
	}
	sql := db.Table(table.Programme + " p")
	sql.LeftJoin(table.Department+" qd", "p.query_department_id=qd.id")
	sql.LeftJoin(table.Department+" qp", "p.query_point_id=qp.id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("p.delete=?", 0)
	sql.And("p.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&programme)
	}
	return programme, err
}

func Del(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	tx, err := db.Begin()
	if err == nil {
		r, _ := tx.Table(table.Programme).Where("id=? AND state=?", id, state.Draft).Data(g.Map{"delete": 1}).Update()
		rows, _ = r.RowsAffected()
	}
	if err == nil && rows > 0 {
		_, _ = delBasis(tx, id)
		_, _ = delBusiness(tx, id)
		_, _ = delContent(tx, id)
		_, _ = delEmphases(tx, id)
		_, _ = delStep(tx, id)
		_, _ = delUser(tx, id)
		_ = tx.Commit()
	} else {
		rows = 0
		_ = tx.Rollback()
	}
	return int(rows), err
}
