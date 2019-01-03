package db_rectifyReport

import (
	"auditIntegralSys/_public/table"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addContent(tx gdb.TX, rectifyReportId int, data g.List) (int, error) {
	for _, v := range data {
		v["rectify_report_id"] = rectifyReportId
	}
	fmt.Println(data)
	if len(data) == 0 {
		return 0, nil
	}
	r, err := tx.BatchInsert(table.RectifyReportContent, data, 5)
	id, _ := r.LastInsertId()
	return int(id), err
}

func delContent(tx gdb.TX, rectifyReportId int) (int, error) {
	r, err := tx.Table(table.RectifyReportContent).Data(g.Map{"delete": 1}).Where("rectify_report_id=?", rectifyReportId).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func GetContents(rectifyReportId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.RectifyReportContent + " rrc")
	sql.LeftJoin(table.DraftContent+" dc", "rrc.draft_content_id=dc.id")
	sql.Fields("rrc.*")
	sql.OrderBy("dc.order asc")
	sql.Where("rrc.delete=? AND rectify_report_id=?", 0, rectifyReportId)
	r, err := sql.All()
	return r.ToList(), err
}
