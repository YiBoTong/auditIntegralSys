package db_integral

import (
	"auditIntegralSys/Audit/entity"
	"auditIntegralSys/Audit/fun"
	entity2 "auditIntegralSys/Org/entity"
	"auditIntegralSys/_public/state"
	"auditIntegralSys/_public/table"
	"auditIntegralSys/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
	"time"
)

func getListSql(db gdb.DB, authorInfo entity2.User, where g.Map) *gdb.Model {
	sql := db.Table(table.Integral + " i")
	sql.LeftJoin(table.Draft+" d", "i.draft_id=d.id")
	sql.LeftJoin(table.Department+" dd", "d.department_id=dd.id")
	sql.LeftJoin(table.Department+" dq", "d.query_department_id=dq.id")
	sql.LeftJoin(table.User+" u", "i.user_id=u.user_id")
	sql.LeftJoin(table.User+" uc", "i.cognizance_user_id=uc.user_id")
	sql.LeftJoin(table.IntegralEdit+" ie", "ie.integral_id=i.id AND ie.delete=0")

	sql.Where("d.delete=?", 0)
	sql.GroupBy("i.id")

	sql = fun.CheckIsMyData(*sql, authorInfo, where)

	if len(where) > 0 {
		sql.And(where)
	}
	return sql
}

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
	sql.Fields("SUM(score) AS sum_score")
	res, err := sql.All()
	if len(res) > 0 {
		for _, v := range res[0] {
			return v.Int(), nil
		}
	}
	return 0, err
}

func GetBetweenTimeScore(userId int, startDay, endDay string) (int, error) {
	db := g.DB()
	sumScore := 0
	sql := db.Table(table.Integral).Where("user_id=? AND `delete`=?", userId, 0)
	sql.And("time BETWEEN ? AND ?", startDay, endDay)
	sql.Fields("SUM(score) AS sum_score")
	res, err := sql.All()
	if len(res) > 0 {
		for _, v := range res {
			sumScore += v["sum_score"].Int()
		}
	}
	return sumScore, err
}

func AddScore(tx gdb.TX, data g.Map) (int, error) {
	r, err := tx.Table(table.Integral).Data(data).Insert()
	row, _ := r.RowsAffected()
	return int(row), err
}

func Count(authorInfo entity2.User, where g.Map) (int, error) {
	db := g.DB()
	r, err := getListSql(db, authorInfo, where).Count()
	return r, err
}

func List(authorInfo entity2.User, offset, limit int, where g.Map) (g.List, error) {
	db := g.DB()
	sql := getListSql(db, authorInfo, where)

	fields := []string{
		"d.*",
		"i.*",
		"ie.state",
		"ie.id as integral_edit_id",
		"u.user_name",
		"uc.user_name as cognizance_user_name",
		"dd.name as department_name",
		"dq.name as query_department_name",
	}
	sql.Fields(strings.Join(fields, ","))
	r, err := sql.Limit(offset, limit).OrderBy("d.id desc").Select()
	return r.ToList(), err
}

func Get(id int, where ...g.Map) (entity.IntegralItem, error) {
	db := g.DB()
	sql := db.Table(table.Integral + " i")
	sql.LeftJoin(table.User+" u", "i.user_id=u.user_id")
	sql.LeftJoin(table.User+" uc", "i.cognizance_user_id=uc.user_id")
	sql.Fields("i.*,u.user_name,uc.user_name as cognizance_user_name")
	sql.Where("i.delete=? AND i.id=?", 0, id)
	if len(where) > 0 {
		sql.And(where[0])
	}
	integralItem := entity.IntegralItem{}
	r, err := sql.One()
	_ = r.ToStruct(&integralItem)
	return integralItem, err
}

func update(tx gdb.TX, id int, data g.Map) (int, error) {
	r, err := tx.Table(table.Integral).Data(data).Where("id=? AND `delete`=0", id).Update()
	row, _ := r.RowsAffected()
	return int(row), err
}

func ChangeScore(id int, data g.Map) (int, error) {
	changeId := 0
	db := g.DB()
	tx, err := db.Begin()
	data["integral_id"] = id
	if err == nil {
		_, _ = delChangeScore(*tx, id)
		changeId, err = addChangeScore(*tx, data)
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return changeId, err
}

func AdoptChangeScore(changeScoreId int, stateStr, suggestion string) (int, error) {
	row := 0
	db := g.DB()
	integralChangeScore := entity.IntegralChangeScore{}
	tx, err := db.Begin()
	if err == nil {
		row, err = updateChangeScore(*tx, changeScoreId,
			g.Map{"state": stateStr, "suggestion": suggestion, "update_time": util.GetLocalNowTimeStr()},
			g.Map{"state": state.Report},
		)
	}
	if row != 0 && err == nil && stateStr == state.Adopt {
		integralChangeScore, err = GetChange(changeScoreId)
	}
	if integralChangeScore.Id != 0 && err == nil {
		row, err = update(*tx, integralChangeScore.IntegralId, g.Map{
			"score": integralChangeScore.Score,
			"time":  integralChangeScore.UpdateTime,
		})
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return row, err
}

// 统计一个整改通知违规人员的违规数量
func GetScoreTotalUserByDraft(draftId int) (g.List, error) {
	db := g.DB()
	fields := []string{
		"i.id",
		"u.user_id",
		"u.user_name",
		"SUM(i.score) AS `score`",
		"SUM(i.money) AS `money`",
	}
	sql := db.Table(table.Integral + " i")
	sql.LeftJoin(table.User+" u", "i.user_id=u.user_id")
	sql.Where("i.draft_id=?", draftId)
	sql.Fields(strings.Join(fields, ","))
	sql.GroupBy("i.id")
	r, err := sql.All()
	return r.ToList(), err
}

func CountMonthScore(startDay, endDay string) (int, error) {
	db := g.DB()
	sql := db.Table(table.Integral).Where("`delete`=?", 0)
	sql.And("time BETWEEN ? AND ?", startDay, endDay)
	sql.Fields("SUM(score) AS sum_score")
	res, err := sql.One()
	return res["sum_score"].Int(), err
}
