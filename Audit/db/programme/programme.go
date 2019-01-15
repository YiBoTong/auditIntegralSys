package db_programme

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
	"time"
)

func add(tx *gdb.TX, data g.Map) (int, error) {
	programme, _ := GetLastOne()
	year := time.Now().Year()
	number := fun.CreateNumber(programme.Year, programme.Number)
	data["year"] = year
	data["number"] = number
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

func getListSql(db gdb.DB, authorId int, where g.Map) *gdb.Model {
	sql := db.Table(table.Programme + " p")
	sql.Where("p.delete=?", 0)
	// 查询自己和别人已发布的数据
	stateStr := state.Publish
	if where["state"] != nil && where["state"] != "" {
		stateStr = gconv.String(where["state"])
		sql.And("(p.author_id=? OR (p.author_id!=? AND (p.state=? OR p.state=?)))", authorId, authorId, stateStr, state.Publish)
		delete(where, "state")
	} else {
		sql.And("(p.author_id=? OR (p.author_id!=? AND p.state=?))", authorId, authorId, state.Publish)
	}
	// 标题模糊查询
	if where["title"] != nil && where["title"] != "" {
		// 查询自己和别人已发布的数据
		likeVal := strings.Replace("%?%", "?", gconv.String(where["title"]), 1)
		sql.And("p.title like ?", likeVal)
		delete(where, "title")
	}
	// 时间区间
	if where["start_time"] != nil && where["start_time"] != "" {
		// 今天开始
		sql.And("(p.start_time < ? OR p.start_time = ?)", where["start_time"], where["start_time"])
		delete(where, "start_time")
	}
	// 时间区间
	if where["end_time"] != nil && where["end_time"] != "" {
		// 今天结束
		sql.And("(p.end_time > ? OR p.end_time = ?)", where["end_time"], where["end_time"])
		delete(where, "end_time")
	}
	sql.GroupBy("p.id")
	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(authorId int, where g.Map) (int, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
	r, err := sql.Count()
	return r, err
}

func List(authorId, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, authorId, where)
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

func GetLastOne() (entity.ProgrammeItem, error) {
	db := g.DB()
	confirmation := entity.ProgrammeItem{}
	sql := db.Table(table.Programme).Where("`delete`=?", 0)
	sql.OrderBy("id desc")
	r, err := sql.One()
	_ = r.ToStruct(&confirmation)
	return confirmation, err
}
