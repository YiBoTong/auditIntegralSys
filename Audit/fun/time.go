package fun

import (
	"fmt"
	"gitee.com/johng/gf/g/util/gconv"
	"strings"
)

func GetBetweenOneMonth(year, month int) (string, string) {
	nextYear := year
	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		nextYear = nextYear + 1
	}
	monthArr := []string{"0", gconv.String(month)}
	nextMonthArr := []string{"0", gconv.String(nextMonth)}
	if month > 9 {
		monthArr = monthArr[1:2]
	}
	if nextMonth > 9 {
		nextMonthArr = nextMonthArr[1:2]
	}
	begin := fmt.Sprint(year, "-", strings.Join(monthArr, ""), "-", "01")
	end := fmt.Sprint(nextYear, "-", strings.Join(nextMonthArr, ""), "-", "01")
	return begin, end
}
