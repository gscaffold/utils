package utils

import (
	"strconv"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// 秒级时间戳字符串转换为时间结构体
func UnixStrToTime(e string) (time.Time, error) {
	data, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return time.Time{}, err //nolint:wrapcheck
	}
	dateTime := time.Unix(data, 0).Local()
	return dateTime, nil
}

// 毫秒级时间戳字符串转换为时间结构体
func UnixMillStrToTime(e string) (time.Time, error) {
	data, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return time.Time{}, err //nolint:wrapcheck
	}
	dateTime := time.Unix(data/1000, 0).Local()
	return dateTime, nil
}

// GetRelativeTime 将一个秒级时间戳转换为相对当前时间的友好字符串表示。
// 例如："刚刚"、"5分钟前"、"3天前"、"2个月前"等。
//
// 参数:
//
//	seconds: 秒级时间戳（Unix 时间戳）
//
// 返回值:
//
//	表示相对时间的字符串。例如：
//	  - 小于1分钟：返回 "刚刚"
//	  - 几分钟内：返回 "X分钟前"
//	  - 几小时内：返回 "X小时前"
//	  - 几天内：返回 "X天前"
//	  - 几周内：返回 "X周前"
//	  - 几个月内：返回 "X个月前"
//	  - 一年以上：返回 "X年前"
func GetRelativeTime(seconds int64) string {
	timestamp := time.Unix(seconds, 0)
	now := time.Now()
	diff := now.Sub(timestamp)

	switch {
	case diff < time.Minute:
		return "刚刚"
	case diff < time.Hour:
		return strconv.Itoa(int(diff.Minutes())) + "分钟前"
	case diff < 24*time.Hour:
		return strconv.Itoa(int(diff.Hours())) + "小时前"
	case diff < 30*24*time.Hour:
		return strconv.Itoa(int(diff.Hours()/24)) + "天前"
	case diff < 365*24*time.Hour:
		return strconv.Itoa(int(diff.Hours()/24/7)) + "周前"
	case diff < 10*365*24*time.Hour:
		return strconv.Itoa(int(diff.Hours()/24/30)) + "个月前"
	default:
		return strconv.Itoa(int(diff.Hours()/24/365)) + "年前"
	}
}
