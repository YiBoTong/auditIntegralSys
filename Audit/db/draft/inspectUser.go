package db_draft

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addInspectUser(tx *gdb.TX, draftId int, userIds string) (int, error) {
	userIdArr := strings.Split(userIds, ",")
	list := []g.Map{}
	for _, v := range userIdArr {
		userId := gconv.Int(v)
		if userId != 0 {
			list = append(list, g.Map{"draft_id": draftId, "user_id": userId})
		}
	}
	if len(list) == 0 {
		return 0, nil
	}
	res, err := tx.BatchInsert(config.DraftInspectUserTbName, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func updateInspectUser(tx *gdb.TX, draftId int, userIds string) (int, error) {
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
	sql := tx.Table(config.DraftInspectUserTbName).Data(g.Map{"delete": 0})
	sql.Where("draft_id=?", draftId)
	sql.And("user_id IN ?", userId)
	res, err := sql.Update()
	row, _ := res.RowsAffected()
	return int(row), err
}

func delInspectUser(tx *gdb.TX, draftId int) (int, error) {
	r, err := tx.Table(config.DraftInspectUserTbName).Where("draft_id=?", draftId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetInspectUser(draftId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.DraftInspectUserTbName + " d")
	sql.LeftJoin(config.UserTbName+" u", "d.user_id=u.user_id")
	sql.Fields("d.*,u.user_name")
	sql.Where("d.draft_id=?", draftId)
	sql.And("d.delete=?", 0)
	sql.OrderBy("d.id asc")
	res, err := sql.All()
	return res.ToList(), err
}
