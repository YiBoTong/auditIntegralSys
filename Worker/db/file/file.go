package db_file

import (
	"auditIntegralSys/_public/config"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func AddFile(fileInfo g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(config.FileTbName, fileInfo)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateFileByIds(tbName, fileIds string, tbId int, tx ...*gdb.TX) (int, error) {
	var rows int = 0
	var row int = 0
	var err error = nil
	fileIdArr := strings.Split(fileIds, ",")
	for _, v := range fileIdArr {
		fId := gconv.Int(v)
		if fId != 0 {
			row, _ = UpdateFile(fId, g.Map{"form": tbName, "form_id": tbId, "delete": 0}, tx[0])
			rows += row
		}
	}
	return rows, err
}

func UpdateFile(fileId int, update g.Map, tx ...*gdb.TX) (int, error) {
	var rows int64 = 0
	var r sql.Result
	var err error = nil
	if len(tx) > 0 {
		r, err = tx[0].Table(config.FileTbName).Where("id=?", fileId).Data(update).Update()
	} else {
		db := g.DB()
		r, err = db.Table(config.FileTbName).Where("id=?", fileId).Data(update).Update()
	}
	rows, _ = r.RowsAffected()
	return int(rows), err
}

func GetFilesByFrom(formId int, form string) (g.List, error) {
	db := g.DB()
	sql := db.Table(config.FileTbName).Where(g.Map{"`delete`": 0})
	sql.And("form_id=?", formId)
	sql.And("`form`=?", form)
	r, err := sql.OrderBy("id asc").All()
	return r.ToList(), err
}

func DelFilesByFrom(formId int, form string, tx ...*gdb.TX) (int, error) {
	var rows int64 = 0
	db := g.DB()
	sql := db.Table(config.FileTbName).Data(g.Map{
		"delete": 1,
	})
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
	sql := tx.Table(config.FileTbName).Data(g.Map{
		"delete": 1,
	})
	sql.Where("form_id=?", formId)
	sql.And("form=?", form)
	r, err := sql.Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}
