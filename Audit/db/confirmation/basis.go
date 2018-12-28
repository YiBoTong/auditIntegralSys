package db_confirmation

import (
	"auditIntegralSys/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func addBasis(tx *gdb.TX, confirmationId int, confirmationBasisIds string) (int, error) {
	confirmationBasisIdsArr := strings.Split(confirmationBasisIds, ",")
	list := []g.Map{}
	for _, v := range confirmationBasisIdsArr {
		confirmationBasisId := gconv.Int(v)
		if confirmationBasisId != 0 {
			list = append(list, g.Map{"confirmation_id": confirmationId, "basis_id": confirmationBasisId})
		}
	}
	if len(list) == 0 {
		return 0, nil
	}
	res, err := tx.BatchInsert(table.ConfirmationBasis, list, 5)
	rows, _ := res.RowsAffected()
	return int(rows), err
}

func EditBasis(confirmationId int, confirmationBasisIds string) (int, error) {
	row := 0
	rows := 0
	db := g.DB()
	tx, err := db.Begin()
	if err == nil {
		row, err = delBasis(tx, confirmationId)
		rows += row
	}
	if err == nil {
		row, err = addBasis(tx, confirmationId, confirmationBasisIds)
		rows += row
	}
	if err == nil {
		_ = tx.Commit()
	} else {
		rows = 0
		_ = tx.Rollback()
	}
	return rows, err
}

func delBasis(tx *gdb.TX, confirmationId int) (int, error) {
	r, err := tx.Table(table.ConfirmationBasis).Where("confirmation_id=?", confirmationId).Data(g.Map{"delete": 1}).Update()
	rows, _ := r.RowsAffected()
	return int(rows), err
}

func GetBasis(confirmationId int) (g.List, error) {
	db := g.DB()
	sql := db.Table(table.ConfirmationBasis + " cb")
	sql.LeftJoin(table.ProgrammeBasis+" pb", "cb.basis_id=pb.id")
	sql.Fields("pb.*")
	sql.Where("cb.confirmation_id=?", confirmationId)
	sql.And("cb.delete=?", 0)
	sql.OrderBy("cb.id asc")
	res, err := sql.All()
	return res.ToList(), err
}
