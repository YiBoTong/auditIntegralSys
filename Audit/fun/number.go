package fun

import (
	"time"
)

// 根据年份生成七位数的编号
// 新的一年编号重新生成
func CreateNumber(preYear, preNumber int) int {
	year := time.Now().Year()
	// 新的一年编号重新生成
	number := 1
	if preNumber != 0 && preYear == year {
		// 相同年份自增1
		number = preNumber + 1
	}
	return number
}
