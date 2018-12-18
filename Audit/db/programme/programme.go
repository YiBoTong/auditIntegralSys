package db_programme

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func add(ctx *gdb.TX, data g.Map) (int, error) {
	var id int64 = 0
	res, err := ctx.Table(config.ProgrammeTbName).Data(data).Insert()
	if err == nil {
		id, _ = res.LastInsertId()
	}
	return int(id), err
}

func edit(ctx *gdb.TX, id int, data g.Map) (int64, error) {
	var row int64 = 0
	res, err := ctx.Table(config.ProgrammeTbName).Data(data).Where("id=?", id).Update()
	if err == nil {
		row, _ = res.RowsAffected()
	}
	return row, err
}

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.ProgrammeTbName + " p")
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
	ctx, err := db.Begin()
	if err == nil {
		id, err = add(ctx, programme)
	}
	if err == nil && id != 0 {
		_, err = addBasis(*ctx, id, basis)
	}
	if err == nil && id != 0 {
		_, err = addContent(ctx, id, content)
	}
	if err == nil && id != 0 {
		_, err = addStep(ctx, id, step)
	}
	if err == nil && id != 0 {
		_, err = addBusiness(ctx, id, business)
	}
	if err == nil && id != 0 {
		_, err = addEmphases(ctx, id, emphases)
	}
	if err == nil && id != 0 {
		_, err = addUser(ctx, id, user)
	}
	if err == nil {
		err = ctx.Commit()
	} else {
		id = 0
		err = ctx.Rollback()
	}
	return id, err
}

func Edit(id int, programme g.Map, basis, content, step, business, emphases, user [][]g.Map) (int, error) {
	db := g.DB()
	var rows int64 = 0
	var row int64 = 0
	ctx, err := db.Begin()
	if err == nil {
		row, err = edit(ctx, id, programme)
		rows += row
	}
	if err == nil && id != 0 {
		//_, _ = delBasis(ctx, id)
		_, err = addBasis(*ctx, id, basis[1])
	}
	if err == nil && id != 0 {
		_, err = addContent(ctx, id, content[1])
	}
	if err == nil && id != 0 {
		_, err = addStep(ctx, id, step[1])
	}
	if err == nil && id != 0 {
		_, err = addBusiness(ctx, id, business[1])
	}
	if err == nil && id != 0 {
		_, err = addEmphases(ctx, id, emphases[1])
	}
	if err == nil && id != 0 {
		_, err = addUser(ctx, id, user[1])
	}
	if err == nil {
		err = ctx.Commit()
	} else {
		id = 0
		err = ctx.Rollback()
	}
	return id, err
}

func Get(id int) (entity.ProgrammeItem, error) {
	db := g.DB()
	programme := entity.ProgrammeItem{}
	fields := []string{
		"p.*",
		"qd.name as query_department_name",
		"qp.name as query_point_name",
		"ud.user_name as det_user_name",
		"ua.user_name as admin_user_name",
	}
	sql := db.Table(config.ProgrammeTbName + " p")
	sql.LeftJoin(config.DepartmentTbName+" qd", "p.query_department_id=qd.id")
	sql.LeftJoin(config.DepartmentTbName+" qp", "p.query_point_id=qp.id")
	sql.LeftJoin(config.UserTbName+" ud", "p.det_user_id=ud.user_id")
	sql.LeftJoin(config.UserTbName+" ua", "p.admin_user_id=ua.user_id")
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
	var row int64 = 0
	ctx, err := db.Begin()
	if err == nil {
		r, _ := ctx.Table(config.ProgrammeTbName).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
		_ = ctx.Rollback()
		row, err = r.RowsAffected()
		rows += row
		//row, err = delBasis(ctx, id)
		//rows += row
	}
	//if err == nil {
	//	_ = ctx.Rollback()
	//	//_, _ = delBasis(id)
	//	//_, _ = delBusiness(id)
	//	//_, _ = delContent(id)
	//	//_, _ = delEmphases(id)
	//	//_, _ = delStep(id)
	//	//_, _ = delUser(id)
	//	//rows, _ = r.RowsAffected()
	//} else {
	//	_ = ctx.Rollback()
	//	rows = 0
	//}
	return int(rows), err
}
