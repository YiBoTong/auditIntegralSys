package db_rectifyReport

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addRectifyReportFiles(tx gdb.TX, rectifyReportId int, fileIds string) (int, error) {
	addIds := strings.Split(fileIds, ",")
	var add []g.Map
	if len(addIds) > 0 && addIds[0] != "" {
		for _, id := range addIds {
			fId := gconv.Int(id)
			if fId > 0 {
				add = append(add, g.Map{
					"rectify_report_id": rectifyReportId,
					"file_id":           fId,
				})
			}
		}
	}
	if len(add) > 0 {
		r, err := tx.BatchInsert(table.RectifyReportFile, add, 5)
		lastId, err := r.LastInsertId()
		return int(lastId), err
	}
	return 0, nil
}

func delRectifyReportFile(tx gdb.TX, rectifyReportId int) (int, error) {
	r, err := tx.Table(table.RectifyReportFile).Where("`delete`=0 AND rectify_report_id=?", rectifyReportId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}
