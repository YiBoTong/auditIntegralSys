package db_notice

import (
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
)

func AddNoticeInform(add []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(config.NoticeInformTbName, add, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func GetNoticeInform(noticeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.NoticeInformTbName + " ni")
	sql.LeftJoin(config.DepartmentTbName+" d", "ni.department_id=d.id")
	sql.Fields("d.*,ni.id as nid")
	sql.Where("ni.notice_id=?", noticeId)
	sql.And("d.delete=?", 0)
	r, err := sql.OrderBy("ni.id desc").Select()
	return r.ToList(), err
}

func DelNoticeInform(noticeId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.NoticeInformTbName).Where("notice_id=?", noticeId).Delete()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}