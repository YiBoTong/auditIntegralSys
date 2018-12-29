package db_integral

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"time"
)

func GetSumScore(punishNoticeId, userId int) (int, error) {
	db := g.DB()
	nowYear := time.Now().Year()
	sql := db.Table(table.Integral)
	// 排除指定的通知书对应的分数
	sql.Where("punish_notice_id!=?", punishNoticeId)
	sql.And("user_id=?", userId)
	sql.And("`delete`=?", 0)
	// 统计一年的分数
	sql.And("time BETWEEN ? AND ?", gconv.String(nowYear)+"-01-01", gconv.String(nowYear+1)+"-01-01")
	sql.Fields("*")
	res, err := sql.All()
	if len(res) > 0 {
		for _, v := range res[0] {
			return v.Int(), nil
		}
	}
	return 0, err
}

func AddScore(tx gdb.TX, data g.Map) (int, error) {
	r, err := tx.Table(table.Integral).Data(data).Insert()
	row, _ := r.RowsAffected()
	return int(row), err
}
