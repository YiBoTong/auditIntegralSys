package db_draft

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addAdminUser(tx *gdb.TX, draftId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	l := len(userIdArr)
	if l == 0 {
		return 0, nil
	}
	list := []g.Map{}
	for i := 0; i < l; i++ {
		list = append(list, g.Map{"draft_id": draftId, "user_id": userIdArr[i]})
	}
	res, err := tx.BatchInsert(config.DraftAdminUserTbName, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateAdminUser(tx *gdb.TX, draftId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	l := len(userIdArr)
	if l == 0 {
		return 0, nil
	}
	userId := g.Slice{}
	for i := 0; i < l; i++ {
		userId = append(userId, gconv.Int(userIdArr[i]))
	}
	sql := tx.Table(config.DraftAdminUserTbName).Data(g.Map{"delete": 0})
	sql.Where("draft_id=?", draftId)
	sql.And("user_id IN ?", userId)
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return int(row), err
}

func delAdminUser(tx *gdb.TX, draftId int) (int, error) {
	r, err := tx.Table(config.DraftAdminUserTbName).Where("draft_id=?", draftId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetAdminUser(draftId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DraftAdminUserTbName + " d")
	sql.LeftJoin(config.UserTbName+" u", "d.user_id=u.user_id")
	sql.Fields("d.*,u.user_name")
	sql.Where("d.draft_id=?", draftId)
	sql.And("d.delete=?", 0)
	sql.OrderBy("d.id asc")
	res, err := sql.All()
	return res.ToList(), err
}
