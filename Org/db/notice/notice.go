package db_notice

import (
	"auditIntegralSys/Org/entity"
	"auditIntegralSys/Worker/db/file"
	"auditIntegralSys/_public/config"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func GetNoticeCount(where g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(config.NoticeTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Count()
	return r, err
}

func GetNotices(offset int, limit int, where g.Map) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.NoticeTbName).Where("`delete`=?", 0)
	if len(where) > 0 {
		sql.And(where)
	}
	r, err := sql.Limit(offset, limit).OrderBy("id desc").Select()
	return r.ToList(), err
}

func GetNotice(id int) (entity.Notice, error) {
	var Notice entity.Notice
	db := g.DB()
	r, err := db.Table(config.NoticeTbName).Where("id=?", id).And("`delete`=?", 0).One()
	_ = r.ToStruct(&Notice)
	return Notice, err
}

func AddNotice(Notice g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	r, err := db.Insert(config.NoticeTbName, Notice)
	if err == nil {
		lastId, err = r.LastInsertId()
	}
	return int(lastId), err
}

func UpdateNotice(id int, Notice g.Map) (int, error) {
	var rows int64 = 0
	db := g.DB()
	r, err := db.Table(config.NoticeTbName).Where("id=?", id).Data(Notice).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

func DelNotice(id int) (int, error) {
	db := g.DB()
	var rows int64 = 0
	r, err := db.Table(config.NoticeTbName).Where("id=?", id).Data(g.Map{"delete": 1}).Update()
	if err == nil {
		rows, _ = r.RowsAffected()
	}
	return int(rows), err
}

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

func GetNoticeFile(noticeId int) ([]map[string]interface{}, error) {
	db := g.DB()
	sql := db.Table(config.NoticeFileTbName + " ni")
	sql.LeftJoin(config.FileTbName+" f", "ni.file_id=f.id")
	sql.Fields("f.*,ni.id as nid")
	sql.Where("ni.notice_id=?", noticeId)
	sql.And("f.delete=?", 0)
	r, err := sql.OrderBy("ni.id desc").Select()
	return r.ToList(), err
}

func AddNoticeFile(add []g.Map) (int, error) {
	var lastId int64 = 0
	db := g.DB()
	// 批次5条数据写入
	r, err := db.BatchInsert(config.NoticeFileTbName, add, 5)
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
					"form":    config.NoticeTbName,
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
	r, err := db.Table(config.NoticeFileTbName).Where("notice_id=?", noticeId).Delete()
	if err == nil {
		rows, _ = r.RowsAffected()
		_, _ = db_file.DelFilesByFrom(noticeId, config.NoticeTbName)
	}
	return int(rows), err
}
