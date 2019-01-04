package db_auditReport

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func addPlan(tx gdb.TX, auditReportId int, content string) (int, error) {
	res, err := tx.Table(table.AuditReportPlan).Data(g.Map{"audit_report_id": auditReportId, "content": content}).Insert()
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func delPlan(tx gdb.TX, auditReportId int) (int, error) {
	r, err := tx.Table(table.AuditReportPlan).Where("audit_report_id=?", auditReportId).Delete()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetPlan(auditReportId int) (entity.AuditReportContent, error) {
	db := g.DB()
	sql := db.Table(table.AuditReportPlan)
	sql.Where("audit_report_id=? AND `delete`=?", auditReportId, 0)
	res, err := sql.One()
	auditReportContent := entity.AuditReportContent{}
	_ = res.ToStruct(&auditReportContent)
	return auditReportContent, err
}
