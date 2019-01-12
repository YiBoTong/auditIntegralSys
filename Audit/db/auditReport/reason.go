package db_auditReport

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addReason(tx gdb.TX, auditReportId int, content string) (int, error) {
	res, err := tx.Table(table.AuditReportReason).Data(g.Map{"audit_report_id": auditReportId, "content": content}).Insert()
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func delReason(tx gdb.TX, auditReportId int) (int, error) {
	r, err := tx.Table(table.AuditReportReason).Where("audit_report_id=?", auditReportId).Delete()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetReason(auditReportId int) (entity.AuditReportContent, error) {
	db := g.DB()
	sql := db.Table(table.AuditReportReason)
	sql.Where("audit_report_id=? AND `delete`=?", auditReportId, 0)
	res, err := sql.One()
	auditReportContent := entity.AuditReportContent{}
	_ = res.ToStruct(&auditReportContent)
	return auditReportContent, err
}
