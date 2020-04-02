package util

import (
	"time"
)

// 时间戳转日期 20060102
func TimestampToYMD(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("20060102")
}

func TimestampToDateTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// 获取目标月份第一天的日期
func GetBeginDateOfMonth(targetMonthDate time.Time) time.Time {
	return time.Date(targetMonthDate.Year(), targetMonthDate.Month(), 1, 0, 0, 0, 0, targetMonthDate.Location())
}

// 获取目标月份最后一天的日期
// lhlyu 9.30 改成获取两个月
// hertz 10.28 修正回一个月
func GetEndDateOfMonth(targetMonthDate time.Time) time.Time {
	return GetBeginDateOfMonth(targetMonthDate).AddDate(0, 1, -1)
}

// 格式化时间对象成日期: yyyy-MM-dd
func FormatToDate(timeObj time.Time) string {
	return timeObj.Format("2006-01-02")
}

// 格式化时间对象成日期: yyyy-MM
func FormatToMonth(timeObj time.Time) string {
	return timeObj.Format("2006-01")
}

// 获取指定地区、日期的零时
func GetZeroClockTime(targetTime time.Time, location *time.Location) time.Time {
	year, month, day := targetTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

//判断某个时间是否是指定天数后的时间
func IsGreaterThanGiveTime(createTime int64, days int) bool {
	return time.Now().Unix() > time.Unix(createTime, 0).Add(time.Hour*24*2).Unix()
}
