package date

import(
	"time"
)

//GetFistDayOfMonth-获取输入时间所在当前月份的第一天
func GetFistDayOfMonth(date time.Time) time.Time {
	year, mon, _ := date.Date()
	return time.Date(year, mon, 1, 0, 0, 0, 0, time.Local)
}

//GetLastDateOfMonth-获取输入时间所在月份的最后一天
func GetLastDateOfMonth(date time.Time)time.Time{
	year, mon, _ := date.Date()
	return time.Date(year, mon+1, 0, 23, 59, 59, 0, time.Local)
}
