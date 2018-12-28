package db_notice

import (
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func GetNoticeFile(noticeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(table.NoticeFile + " ni")
	sql.LeftJoin(table.File+" f", "ni.file_id=f.id")
	sql.Fields("f.*,ni.id as nid")
	sql.Where("ni.notice_id=?", noticeId)
	sql.And("f.delete=?", 0)
	r, err := sql.OrderBy("ni.id asc").Select()
	return r.ToList(), err
}

func AddNoticeFile(add []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(table.NoticeFile, add, 5)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func AddNoticeFiles(noticeId int, fileIds string) error {
	addIds := strings.Split(fileIds, ",")
	var err error = nil
	if len(addIds) > 0 && addIds[0] != "" {
		var add []g.Map
		for _, id := range addIds {
			fId := gconv.Int(id)
			if fId > 0 {
				add = append(add, g.Map{
					"notice_id": noticeId,
					"file_id":   fId,
				})
				_, err = db_file.UpdateFile(fId, g.Map{
					"form":    table.Notice,
					"form_id": noticeId,
					"delete":  0,
				})
				if err != nil {
					break
				}
			}
		}
		if err == nil {
			_, err = AddNoticeFile(add)
		}
	}
	return err
}

func DelNoticeFile(noticeId int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(table.NoticeFile).Where("notice_id=?", noticeId).Delete()
	if err == nil {
		rows, _ = r.RowsAffected()
		_, _ = db_file.DelFilesByFrom(noticeId, table.Notice)
	}
	return int(rows), err
}