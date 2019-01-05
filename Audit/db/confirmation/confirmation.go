package db_confirmation

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/punishNotice"
	"auditIntegralSys/Audit/db/rectify"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
)

func getInspectUser(draftId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.DraftInspectUser + " d")
	sql.Where("d.draft_id=?", draftId)
	sql.And("d.delete=?", 0)
	sql.OrderBy("d.id asc")
	res, err := sql.All()
	return res.ToList(), err
}

func Count(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Confirmation + " c")
	sql.LeftJoin(table.Draft+" d", "c.draft_id=d.id")
	sql.Where("c.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func List(offset int, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Confirmation + " c")
	sql.LeftJoin(table.Draft+" d", "c.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.Fields("d.*,c.*,dd.name as department_name,dq.name as query_department_name,p.title as programme_title")
	sql.Where("c.delete=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("c.id desc").Select()
	return r.ToList(), err
}

func Add(tx gdb.TX, draftId int) (int, error) {
	r, err := tx.Table(table.Confirmation).Data(g.Map{"draft_id": draftId}).Insert()
	id, _ := r.LastInsertId()
	return int(id), err
}

func Get(id int, where ...g.Map) (entity.ConfirmationItem, error) {
	db := g.DB()
	confirmation := entity.ConfirmationItem{}
	sql := db.Table(table.Confirmation)
	sql.Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.One()
	if err == nil {
		_ = r.ToStruct(&confirmation)
	}
	return confirmation, err
}

func Update(id int, data g.Map, where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.Confirmation).Data(data).Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func UpdateTX(tx *gdb.TX,id int, data g.Map, where ...g.Map) (int, error) {
	sql := tx.Table(table.Confirmation).Data(data).Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Publish(id int, TX ...*gdb.TX) (int, error) {
	db := g.DB()
	row := 0
	rows := 0
	confirmation := entity.ConfirmationItem{}
	var err error = nil
	tx := &gdb.TX{}
	if len(TX) == 0 {
		tx, err = db.Begin()
	} else {
		tx = TX[0]
	}
	if err == nil {
		var rowNum int64 = 0
		// 只有草稿的数据才能发布
		r, _ := tx.Table(table.Confirmation).Data(g.Map{
			"state": check.D_publish,
		}).Where("`delete`=? AND state IN (?)", 0, g.Slice{check.D_draft}).And("id=?", id).Update()
		rowNum, _ = r.RowsAffected()
		row = int(rowNum)
		if row > 0 {
			confirmation, _ = Get(id)
		}
	}
	if row != 0 && confirmation.Id != 0 && err == nil {
		// 生成整改通知
		_, err = db_rectify.Add(*tx, confirmation.Id, confirmation.DraftId)
		rows += 1
	}
	if row != 0 && confirmation.Id != 0 && err == nil {
		// 生成惩罚通知书
		userIdArr := []int{}
		userList, _ := getInspectUser(confirmation.DraftId)
		for _, v := range userList {
			userId := gconv.Int(v["user_id"])
			if userId == 0 {
				continue
			}
			// 对每一个被检查人创建一个惩罚通知书
			userIdArr = append(userIdArr, userId)
		}
		_, err = db_punishNotice.Add(*tx, confirmation.Id, confirmation.DraftId, userIdArr)
		rows += 1
	}
	if len(TX) == 0 {
		if err == nil {
			_ = tx.Commit()
		} else {
			row = 0
			_ = tx.Rollback()
		}
	}
	return int(row), err
}
