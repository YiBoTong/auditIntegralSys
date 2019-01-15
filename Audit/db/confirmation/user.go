package db_confirmation

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addUser(tx *gdb.TX, confirmationId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	list := []g.Map{}
	for _, v := range userIdArr {
		userId := gconv.Int(v)
		if userId != 0 {
			list = append(list, g.Map{"confirmation_id": confirmationId, "user_id": userId})
		}
	}
	if len(list) == 0 {
		return 0, nil
	}
	res, err := tx.BatchInsert(table.ConfirmationUser, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateUser(tx *gdb.TX, confirmationId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	userId := g.Slice{}
	for _, v := range userIdArr {
		uId := gconv.Int(v)
		if uId != 0 {
			userId = append(userId, uId)
		}
	}
	if len(userId) == 0 {
		return 0, nil
	}
	sql := tx.Table(table.ConfirmationUser).Data(g.Map{"delete": 0})
	sql.Where("confirmation_id=?", confirmationId)
	sql.And("user_id IN ?", userId)
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return int(row), err
}

func delUser(tx *gdb.TX, confirmationId int) (int, error) {
	r, err := tx.Table(table.ConfirmationUser).Where("confirmation_id=?", confirmationId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetUser(confirmationId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.ConfirmationUser + " c")
	sql.LeftJoin(table.User+" u", "c.user_id=u.user_id")
	sql.Fields("c.*,u.user_name")
	sql.Where("c.confirmation_id=?", confirmationId)
	sql.And("c.delete=?", 0)
	sql.OrderBy("c.id asc")
	res, err := sql.All()
	return res.ToList(), err
}

func GetBetweenTimeNum(userId int, startDay, endDay string) (int, error) {
	db := g.DB()
	sql := db.Table(table.ConfirmationUser + " cu")
	sql.LeftJoin(table.Confirmation+" c", "cu.confirmation_id=c.id")
	sql.LeftJoin(table.Draft+" d", "c.draft_id=d.id")
	sql.Where("cu.user_id=? AND cu.delete=? AND d.delete=?", userId, 0, 0)
	sql.And("d.update_time BETWEEN ? AND ?", startDay, endDay)
	sql.GroupBy("d.id")
	sum, err := sql.Count()
	return sum, err
}
