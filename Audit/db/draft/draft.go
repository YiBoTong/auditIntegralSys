package db_draft

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/auditNotice"
	"auditIntegralSys/Audit/db/confirmation"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
	"time"
)

func add(tx *gdb.TX, data g.Map) (int, error) {
	draft, _ := GetLastOne()
	year := time.Now().Year()
	number := fun.CreateNumber(draft.Year, draft.Number)
	data["year"] = year
	data["number"] = number
	res, err := tx.Table(table.Draft).Data(data).Insert()
	id, _ := res.LastInsertId()
	return int(id), err
}

func edit(tx *gdb.TX, id int, data g.Map, where ...g.Map) (int64, error) {
	sql := tx.Table(table.Draft).Data(data).Where("id=?", id)
	sql.And("`delete`=?", 0)
	if len(where) > 0 {
		for k, v := range where[0] {
			sql.And(k, v)
		}
	}
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return row, err
}

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.Draft + " d")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" qdd", "d.query_department_id=qdd.id")
	sql.LeftJoin(table.Introduction+" i", "d.id=i.draft_id")

	sql.Where("d.delete=?", 0)
	sql.GroupBy("d.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

func Count(authorInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	sql := getListSql(db, authorInfo, where)
	r, err := sql.Count()
	return r, err
}

func List(authorInfo entity2.User, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, authorInfo, where)
	fields := []string{
		"d.*",
		"i.id as introduction_id",
		"p.title as programme_title",
		"dd.name as department_name",
		"qdd.name as query_department_name",
		"i.id as introduction_id",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("d.id desc").Select()
	return r.ToList(), err
}

func Add(draft g.Map, content []g.Map, queryUserLeader int, queryUsers, adminUsers, fileIds string) (int, error) {
	id := 0
	thisYear := time.Now().Year()
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		id, err = add(tx, draft)
	}
	if id != 0 {
		if err == nil {
			_, err = addContent(tx, id, content)
		}
		if err == nil {
			_, err = addAdminUser(tx, id, adminUsers)
		}
		if err == nil {
			_, err = addQueryUser(tx, id, queryUsers, queryUserLeader)
		}
		if err == nil {
			_, err = db_file.UpdateFileByIds(table.Draft, fileIds, id, tx)
		}
		if err == nil && draft["state"] == state.Publish {
			// 生成事实确认书
			_, err = db_confirmation.Add(*tx, id)
			if err == nil {
				// 生成审计通知书
				auditNoticeItem, _ := db_auditNotice.GetLastOne()
				_, err = db_auditNotice.Add(tx, g.Map{
					"draft_id": id,
					"year":     thisYear,
					"number":   fun.CreateNumber(auditNoticeItem.Year, auditNoticeItem.Number),
				})
			}
		}
	}
	if err == nil {
		err = tx.Commit()
	} else {
		id = 0
		err = tx.Rollback()
	}
	return id, err
}

func Edit(id int, draft g.Map, content [2][]g.Map, queryUserLeader int, queryUsers, adminUsers, fileIds string, where ...g.Map) (int, error) {
	row := 0
	rows := 0
	thisYear := time.Now().Year()
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		var r int64 = 0
		r, err = edit(tx, id, draft, where...)
		rows += int(r)
	}
	if err == nil && rows > 0 {
		_, _ = delContent(tx, id)
		row, err = addContent(tx, id, content[0])
		rows += row
		row, err = updateContent(tx, id, content[1])
		rows += row
		if err == nil {
			_, _ = delAdminUser(tx, id)
			row, err = addAdminUser(tx, id, adminUsers)
			rows += row
		}
		if err == nil {
			_, _ = delQueryUser(tx, id)
			row, err = addQueryUser(tx, id, queryUsers, queryUserLeader)
			rows += row
		}
		if err == nil {
			_, _ = db_file.DelFilesByFromTx(id, table.Draft, tx)
			row, err = db_file.UpdateFileByIds(table.Draft, fileIds, id, tx)
			rows += row
		}
		if err == nil && draft["state"] == state.Publish {
			// 生成事实确认书
			_, err = db_confirmation.Add(*tx, id)
			if err == nil {
				// 生成审计通知书
				auditNoticeItem, _ := db_auditNotice.GetLastOne()
				_, err = db_auditNotice.Add(tx, g.Map{
					"draft_id": id,
					"year":     thisYear,
					"number":   fun.CreateNumber(auditNoticeItem.Year, auditNoticeItem.Number),
				})
			}
		}
	}
	if err == nil {
		err = tx.Commit()
	} else {
		rows = 0
		err = tx.Rollback()
	}
	return rows, err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Draft).Data(data).Where("`delete`=?", 0).And("id=?", id)
	if len(where) > 0 {
		for k, v := range where[0] {
			sql.And(k, v)
		}
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Publish(id int) (int, error) {
	db := g.DB()
	row := 0
	rows := 0
	tx, err := db.Begin()
	if err == nil {
		var rowNum int64 = 0
		// 只有草稿的数据才能发布
		r, _ := tx.Table(table.Draft).Data(g.Map{
			"state": check.D_publish,
		}).Where("`delete`=? AND state IN (?)", 0, g.Slice{check.D_draft}).And("id=?", id).Update()
		rowNum, _ = r.RowsAffected()
		row = int(rowNum)
	}
	if row != 0 && err == nil {
		// 生成事实确认书
		_, err = db_confirmation.Add(*tx, id)
		rows += 1
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		row = 0
		_ = tx.Rollback()
	}
	return int(row), err
}

func Get(id int) (entity.DraftItem, error) {
	db := g.DB()
	draft := entity.DraftItem{}
	fields := []string{
		"d.*",
		"p.title as programme_title",
		"qd.name as query_department_name",
		"dm.name as department_name",
	}
	sql := db.Table(table.Draft + " d")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" qd", "d.query_department_id=qd.id")
	sql.LeftJoin(table.Department+" dm", "d.department_id=dm.id")
	sql.Fields(strings.Join(fields, ","))
	sql.Where("d.delete=?", 0)
	sql.And("d.id=?", id)
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&draft)
	}
	return draft, err
}

func Del(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	tx, err := db.Begin()
	if err == nil {
		r, _ := tx.Table(table.Draft).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
		rows, _ = r.RowsAffected()
	}
	if err == nil && rows > 0 {
		_, _ = delContent(tx, id)
		_, _ = delAdminUser(tx, id)
		_, _ = delInspectUser(tx, id)
		_, _ = delQueryUser(tx, id)
		_, _ = delReviewUser(tx, id)
		_, _ = db_file.DelFilesByFromTx(id, table.Draft, tx)
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return int(rows), err
}

func GetLastOne() (entity.DraftItem, error) {
	db := g.DB()
	confirmation := entity.DraftItem{}
	sql := db.Table(table.Draft).Where("`delete`=?", 0)
	sql.OrderBy("id desc")
	r, err := sql.One()
	_ = r.ToStruct(&confirmation)
	return confirmation, err
}
