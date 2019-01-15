package db_clause

import (
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func GetClauseFile(ClauseId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ClauseFile + " cf")
	sql.LeftJoin(table.File+" f", "cf.file_id=f.id")
	sql.Fields("f.*,cf.id as nid")
	sql.Where("cf.clause_id=?", ClauseId)
	sql.And("f.delete=?", 0)
	r, err := sql.OrderBy("cf.id asc").All()
	return r.ToList(), err
}

func AddClauseFile(add []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(table.ClauseFile, add, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func AddClauseFiles(ClauseId int, fileIds string) error {
	addIds := strings.Split(fileIds, ",")
	err := error(nil)
	if len(addIds) > 0 && addIds[0] != "" {
		var add []g.Map
		for _, id := range addIds {
			fId := gconv.Int(id)
			if fId > 0 {
				add = append(add, g.Map{
					"clause_id": ClauseId,
					"file_id":   fId,
				})
				_, err = db_file.UpdateFile(fId, g.Map{
					"form":    table.Clause,
					"form_id": ClauseId,
					"delete":  0,
				})
				if err != nil {
					break
				}
			}
		}
		if err == nil {
			_, err = AddClauseFile(add)
		}
	}
	return err
}

func DelClauseFile(ClauseId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.ClauseFile).Where("clause_id=?", ClauseId).Delete()
	if err == nil {
		rows, _ = r.RowsAffected()
		_, _ = db_file.DelFilesByFrom(ClauseId, table.Clause)
	}
	return int(rows), err
}