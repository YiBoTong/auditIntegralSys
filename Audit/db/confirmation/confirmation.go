package db_confirmation

import (
	"auditIntegralSys/Audit/check"
	"auditIntegralSys/Audit/db/punishNotice"
	"auditIntegralSys/Audit/db/rectify"
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
	"time"
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

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.Confirmation + " c")
	sql.LeftJoin(table.Draft+" d", "c.draft_id=d.id")
	sql.LeftJoin(table.Programme+" p", "d.programme_id=p.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")

	sql.Where("c.delete=?", 0)
	sql.GroupBy("c.id")

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
		"c.*",
		"dd.name as department_name",
		"dq.name as query_department_name",
		"p.title as programme_title",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("c.id desc").Select()
	return r.ToList(), err
}

func Add(tx gdb.TX, draftId int) (int, error) {
	confirmation, _ := GetLastOne()
	year := time.Now().Year()
	number := fun.CreateNumber(confirmation.Year, confirmation.Number)
	r, err := tx.Table(table.Confirmation).Data(g.Map{
		"draft_id": draftId,
		"year":     year,
		"number":   number,
	}).Insert()
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

func UpdateTX(tx *gdb.TX, id int, data g.Map, where ...g.Map) (int, error) {
	sql := tx.Table(table.Confirmation).Data(data).Where("id=? AND `delete`=?", id, 0)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Edit(id, thisUserId int, stateStr string, content [2][]g.Map, users, basisIds, fileIds string, where ...g.Map) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	confirmation := entity.ConfirmationItem{}
	tx, err := db.Begin()
	if err == nil {
		_, err = UpdateTX(tx, id, g.Map{"author_id": thisUserId})
	}
	if err == nil {
		_, _ = delContent(tx, id)
		row, err = addContent(tx, id, content[0])
		rows += row
		row, err = updateContent(tx, id, content[1])
		rows += row
	}

	if err == nil {
		_, _ = delUser(tx, id)
		row, err = addUser(tx, id, users)
		rows += row
	}
	if err == nil {
		_, _ = delBasis(tx, id)
		row, err = addBasis(tx, id, basisIds)
		rows += row
	}
	if err == nil {
		_, _ = db_file.DelFilesByFromTx(id, table.Confirmation, tx)
		row, err = db_file.UpdateFileByIds(table.Confirmation, fileIds, id, tx)
		rows += row
	}
	if err == nil && stateStr == state.Publish { // 发布状态
		confirmation, _ = Get(id)
		// 生成整改通知
		_, err = db_rectify.Add(*tx, id, confirmation.DraftId)
		if confirmation.Id != 0 && err == nil && len(users) > 0 {
			// 生成惩罚通知书
			userIdArr := []int{}
			userList := strings.Split(users, ",")
			for _, v := range userList {
				userId := gconv.Int(v)
				if userId != 0 {
					// 对每一个被检查人创建一个惩罚通知书
					userIdArr = append(userIdArr, userId)
				}
			}
			if len(userIdArr) > 0 {
				_, err = db_punishNotice.Add(*tx, confirmation.Id, confirmation.DraftId, userIdArr)
				rows += 1
			}

		}
		if err == nil {
			_, err = UpdateTX(tx, id, g.Map{"state": stateStr})
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

func Publish(id int, TX ...*gdb.TX) (int, error) {
	db := g.DB()
	row := 0
	rows := 0
	confirmation := entity.ConfirmationItem{}
	err := error(nil)
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

func GetLastOne() (entity.ConfirmationItem, error) {
	db := g.DB()
	confirmation := entity.ConfirmationItem{}
	sql := db.Table(table.Confirmation).Where("`delete`=?", 0)
	sql.OrderBy("id desc")
	r, err := sql.One()
	_ = r.ToStruct(&confirmation)
	return confirmation, err
}

// 获取违规分类最多的事实确认书
func GetViolationTopDepartmentByConfirmationId(num int, confirmationIds g.Slice) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.Confirmation + " c")
	sql.LeftJoin(table.Draft+" d", "c.draft_id=d.id")
	sql.LeftJoin(table.Department+" dm", "d.query_department_id=dm.id")
	sql.Where("c.id IN ? AND d.delete=?", confirmationIds, 0)
	sql.GroupBy("d.query_department_id")
	sql.OrderBy("sum")
	sql.Fields("d.query_department_id,COUNT(1) AS `sum`")
	sql.Limit(0, num)
	res, err := sql.Select()
	return res.ToList(), err
}
