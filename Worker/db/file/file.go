package db_file

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func AddFile(fileInfo g.Map) (int,error) {
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