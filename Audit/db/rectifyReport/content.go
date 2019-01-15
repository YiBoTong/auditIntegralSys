package db_rectifyReport

import (
	"auditIntegralSys/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
)

func addContent(tx gdb.TX, rectifyReportId int, data g.List) (int, error) {
	var id int64 = 0
	var r sql.Result
	err := error(nil)
	for _, v := range data {
		v["rectify_report_id"] = rectifyReportId
	}
	if len(data) == 0 {
		return 0, nil
	}
	for _, v := range data {
		userIds := v["userIds"]
		delete(v, "userIds")
		r, err = tx.Table(table.RectifyReportContent).Data(v).Insert()
		id, _ = r.LastInsertId()
		if err == nil {
			_, err = addContentUser(tx, rectifyReportId, int(id), gconv.String(userIds))
		}
	}
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
