package db_rectifyReport

import (
	"auditIntegralSys/Audit/db/auditReport"
	"auditIntegralSys/Audit/db/draft"
	"auditIntegralSys/Audit/db/rectify"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"errors"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
)

func Get(id int) (entity.RectifyReportItem, error) {
	db := g.DB()
	rectifyReportItem := entity.RectifyReportItem{}
	r, err := db.Table(table.RectifyReport).Where("id=? AND `delete`=?", id, 0).One()
	_ = r.ToStruct(&rectifyReportItem)
	return rectifyReportItem, err
}

func GetByRectifyId(rectifyId int) (entity.RectifyReportItem, error) {
	db := g.DB()
	rectifyReportItem := entity.RectifyReportItem{}
	r, err := db.Table(table.RectifyReport).Where("rectify_id=? AND `delete`=?", rectifyId, 0).One()
	_ = r.ToStruct(&rectifyReportItem)
	return rectifyReportItem, err
}

func add(tx gdb.TX, data g.Map) (int, error) {
	r, err := tx.Table(table.RectifyReport).Data(data).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func update(tx gdb.TX, id int, data g.Map, where ...g.Map) (int, error) {
	sql := tx.Table(table.RectifyReport).Data(data).Where("id=?", id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func del(tx gdb.TX, rectifyId int) (int, error) {
	sql := tx.Table(table.RectifyReport).Data(g.Map{"delete": 1}).Where("rectify_id=?", rectifyId)
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Add(rectifyId int, fileIds string, data g.Map, content g.List) (int, error) {
	db := g.DB()
	id := 0
	tx, err := db.Begin()
	rectifyReportItem, _ := GetByRectifyId(rectifyId)
	if rectifyReportItem.State == state.Publish {
		// 此整改通知已经被填写并上报了
		err = errors.New("has been filled in")
	}
	if err == nil {
		_, _ = del(*tx, rectifyId)
		_, _ = delContent(*tx, rectifyReportItem.Id)
		id, err = add(*tx, data)
	}
	if err == nil {
		_, _ = delContentUser(*tx, id)
		_, err = addContent(*tx, id, content)
	}
	if err == nil {
		_, err = db_file.DelFilesByFromTx(rectifyReportItem.Id, table.RectifyReport, tx)
	}
	if err == nil {
		_, err = db_file.UpdateFileByIds(table.RectifyReport, fileIds, id, tx)
	}
	if err == nil {
		_, _ = delRectifyReportFile(*tx, rectifyReportItem.Id)
		_, err = addRectifyReportFiles(*tx, id, fileIds)
	}
	if err == nil && data["state"] == state.Publish {
		// 生成审计报告
		rectifyItem, _ := db_rectify.Get(rectifyId)
		draftItem, _ := db_draft.Get(rectifyItem.DraftId)
		_, err = db_auditReport.Add(*tx, g.Map{
			"programme_id":      draftItem.ProgrammeId,      // 方案ID
			"draft_id":          rectifyItem.DraftId,        // 工作底稿ID
			"confirmation_id":   rectifyItem.ConfirmationId, // 事实确认书ID
			"rectify_report_id": id,                         // 整改报告ID
		})
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return id, err
}

func Edit(id int, fileIds string, data g.Map, content g.List) (int, error) {
	db := g.DB()
	row := 0
	tx, err := db.Begin()
	if err == nil {
		row, err = update(*tx, id, data)
	}
	if err == nil && row != 0 {
		_, err = delContent(*tx, id)
	}
	if err == nil {
		_, _ = delContentUser(*tx, id)
		row, err = addContent(*tx, id, content)
	}
	if err == nil {
		_, _ = db_file.UpdateFileByIds(table.RectifyReport, fileIds, id, tx)
		_, err = db_file.UpdateFileByIds(table.RectifyReport, fileIds, id, tx)
	}
	if err == nil {
		_, _ = delRectifyReportFile(*tx, id)
		_, err = addRectifyReportFiles(*tx, id, fileIds)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return row, err
}
