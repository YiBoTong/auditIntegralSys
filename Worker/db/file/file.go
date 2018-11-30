package db_file

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
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

func UpdateFile(fileId int, update g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(config.FileTbName).Where("id=?", fileId).Data(update).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelFilesByFrom(formId int, form string) (int, error) {
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
