package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("E:/桌面/新建文件夹/添加人员.xlsx")
	if err == nil {
		rows := xlsx.GetRows("人员导入模版")
		for index, row := range rows {
			if index == 0 {
				continue
			}
			for _, colCell := range row {
				//fmt.Print(colCell, "\t")
				fmt.Println(colCell)
			}
			fmt.Println()
		}
	}
}
