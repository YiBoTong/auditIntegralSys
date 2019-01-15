package db_file

import (
	"auditIntegralSys/Worker/entity"
	"auditIntegralSys/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func AddFile(fileInfo g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(table.File, fileInfo)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateFileByIds(tbName, fileIds string, tbId int, tx ...*gdb.TX) (int, error) {
	var rows int = 0
	var row int = 0
	err := error(nil)
	fileIdArr := strings.Split(fileIds, ",")
	if len(fileIds) > 0 && len(fileIds) > 0 {
		for _, v := range fileIdArr {
			fId := gconv.Int(v)
			if fId != 0 {
				row, _ = UpdateFile(fId, g.Map{"form": tbName, "form_id": tbId, "delete": 0}, tx...)
				rows += row
			}
		}
	}
	return rows, err
}

func UpdateFile(fileId int, update g.Map, tx ...*gdb.TX) (int, error) {
	var rows int64 = 0
	var r sql.Result
	err := error(nil)
	if len(tx) > 0 {
		r, err = tx[0].Table(table.File).Where("id=?", fileId).Data(update).Update()
	} else {
		db := g.DB()
		r, err = db.Table(table.File).Where("id=?", fileId).Data(update).Update()
	}
	rows, _ = r.RowsAffected()
	return int(rows), err
}

func GetFilesByFrom(formId int, form string) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.File).Where(g.Map{"`delete`": 0})
	sql.And("form_id=?", formId)
	sql.And("`form`=?", form)
	r, err := sql.OrderBy("id asc").All()
	return r.ToList(), err
}

func Get(formId int) (entity.File, error) {
	db := g.DB()
	sql := db.Table(table.File).Where(g.Map{"`delete`": 0, "id": formId})
	r, err := sql.One()
	file := entity.File{}
	_ = r.ToStruct(&file)
	return file, err
}

func DelFilesByFrom(formId int, form string) (int, error) {
	var rows int64 = 0
	db := g.DB()
	sql := db.Table(table.File).Data(g.Map{"delete": 1})
	sql.Where("form_id=?", formId)
	sql.And("form=?", form)
	r, err := sql.Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelFilesByFromTx(formId int, form string, tx *gdb.TX) (int, error) {
	var rows int64 = 0
	sql := tx.Table(table.File).Data(g.Map{"delete": 1})
	sql.Where("form_id=?", formId)
	sql.And("form=?", form)
	r, err := sql.Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
