package fun

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadSpreadsheets(path, sheet string) ([][]string, error) {
	var rows [][]string
	xlsx, err := excelize.OpenFile(path)
	if err == nil {
		rows = xlsx.GetRows(sheet)
	}
	return rows, err
}
