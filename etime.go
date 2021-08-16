package etime

import (
	"strings"
	"time"
)

// Today 获取当天的开始和结束时间戳
func Today() (start, end time.Time) {
	return TimeOfDay(time.Now())
}

// Yesterday 获取昨天的开始和结束时间戳
func Yesterday() (start, end time.Time) {
	return DayAgoOrAfter(-1)
}

// Week 获取当天的开始和结束时间戳
func Week() (time.Time, time.Time) {
	t := time.Now()
	return WeekFirst(t), Weekend(t)
}

// Month 获取当月的开始和结束时间戳
func Month() (time.Time, time.Time) {
	t := time.Now()
	return MonthFirst(t), MonthLast(t)
}

// WeekFirst 获取某个时间所在周的第一天
func WeekFirst(t time.Time) time.Time {
	w := int(t.Weekday())
	if w == 0 {
		w = 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()-w+1, 0, 0, 0, 0, t.Location())
}

// Weekend 获取某个时间所在周的周末
func Weekend(t time.Time) time.Time {
	w := int(t.Weekday())
	if w == 0 {
		w = 7
	}
	return time.Date(t.Year(), t.Month(), t.Day()-w+7, 23, 59, 59, 0, t.Location())
}

// MonthFirst 获取某个时间所在月份的第一天
func MonthFirst(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// MonthLast 获取某个时间所在月份的最后一天
func MonthLast(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location())
}

// DayFirst 获取某个时间的0点
func DayFirst(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// DayLast 获取某个时间的结束时间
func DayLast(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// TimeOfDay 获取某天的开始和结束时间
func TimeOfDay(d time.Time) (start, end time.Time) {
	return DayFirst(d), DayLast(d)
}

// DayAgoOrAfter 以当天为准 获取 几天前或者几天后的 开始和结束时间
func DayAgoOrAfter(n int) (start, end time.Time) {
	return DayOfAgoOrAfter(time.Now(), n)
}

// DayOfAgoOrAfter 以某天为准 获取 几天前或者几天后的 开始和结束时间
func DayOfAgoOrAfter(d time.Time, n int) (start, end time.Time) {
	t := d.AddDate(0, 0, n)
	return TimeOfDay(t)
}

// Format Returns a date string of a specific format
// e.g Format(t ,"Y/m/d H:i:s")
func Format(d time.Time, spec ...string) string {
	s := ""
	if len(spec) > 0 {
		s = ParseSpec(spec[0])
	} else {
		s = "2006-01-02 15:04:05"
	}

	return d.Format(s)
}

func ParseSpec(spec string) string {
	return strings.NewReplacer(
		"Y",
		"2006",
		"m",
		"01",
		"d",
		"02",
		"H",
		"15",
		"i",
		"04",
		"s",
		"05",
	).Replace(spec)
}

// Str2Time 返回 日期字符串=>时间对象
// e.g Str2Time('2021-01-03 23/45/46' ,"Y/m/d H:i:s")
func Str2Time(date string, spec ...string) time.Time {
	layout := ParseSpec(spec[0])
	t, _ := time.ParseInLocation(layout, date, time.Local)
	return t
}
