package db_rectifyReport

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addContentUser(tx gdb.TX, rectifyReportId, rectifyReportContentId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	list := []g.Map{}
	for _, v := range userIdArr {
		userId := gconv.Int(v)
		if userId != 0 {
			list = append(list, g.Map{"rectify_report_id": rectifyReportId, "rectify_report_content_id": rectifyReportContentId, "user_id": userId})
		}
	}
	if len(list) == 0 {
		return 0, nil
	}
	res, err := tx.BatchInsert(table.RectifyReportContentUser, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func delContentUser(tx gdb.TX, rectifyReportId int) (int, error) {
	r, err := tx.Table(table.RectifyReportContentUser).Data(g.Map{"delete": 1}).Where("rectify_report_id=?", rectifyReportId).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func GetContentUsers(rectifyReportContentId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.RectifyReportContentUser + " ru")
	sql.LeftJoin(table.User+" u", "ru.user_id=u.user_id")
	sql.Where("ru.delete=? AND ru.rectify_report_content_id=?", 0, rectifyReportContentId)
	sql.Fields("u.*")
	sql.OrderBy("ru.id asc")
	r, err := sql.All()
	return r.ToList(), err
}
