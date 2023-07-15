package utils

import "time"

func ParseDate(dateStr string) time.Time {
	// 解析日期字符串并返回对应的时间对象
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}

func ParseDateTime(dateStr string) time.Time {

	// 解析日期字符串并返回对应的时间对象
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	return t
}
