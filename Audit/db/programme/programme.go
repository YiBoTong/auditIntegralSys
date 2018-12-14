package db_programme

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

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
		_, err = addBasis(ctx, id, basis)
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

func Get(id int) (entity.Programme, error) {
	return entity.Programme{}, nil
}

func add(ctx *gdb.Tx, data g.Map) (int, error) {
	var id int64 = 0
	res, err := ctx.Table(config.ProgrammeTbName).Data(data).Insert()
	if err == nil {
		id, _ = res.LastInsertId()
	}
	return int(id), err
}
