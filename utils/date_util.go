package utils

import (
	"time"
	"math/rand"
)

type _dateUtil struct {

}

var DayUnixTime int64 = 24 * 60 * 60

//time.Utc -8h
var EightHours int64 = 8 * 60 * 60

//yyyy-MM-10 hh:mm:ss
//day 1 return  yyyy-MM-11 hh:mm:ss
func (_dateUtil) AddDay(day int) time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()+day, t.Hour(), t.Minute(), t.Second(), 0, time.Local)
}
func (_dateUtil) AddHour(hour int) time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+hour, t.Minute(), t.Second(), 0, time.Local)
}
func (_dateUtil) AddMinute(minute int) time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+minute, t.Second(), 0, time.Local)
}
//yyyy-MM-dd hh:mm:ss
func (_dateUtil) GetTime(day int) time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day()+day, t.Hour(), t.Minute(), t.Second(), 0, time.Local)
}
func (_dateUtil) GetDayByMinute(minute int)time.Duration  {
	return time.Minute*time.Duration(60*24-rand.Int63n(int64(minute)))
}
func (_dateUtil) GetMonthByMinute(minute int)time.Duration  {
	return time.Minute*time.Duration(60*31*24-rand.Int63n(int64(minute)))
}
func (_dateUtil) GetMonthByDay(day int)time.Duration  {
	return time.Hour*time.Duration(31*24-24*rand.Int63n(int64(day)))
}
func (_dateUtil) GetMonthByHour(hour int)time.Duration  {
	return time.Hour*time.Duration(31*24-rand.Int63n(int64(hour)))
}
//yyyy-10-dd hh:mm:ss
//month 1 return  yyyy-11-dd hh:mm:ss
func (_dateUtil) AddMonth(month int) time.Time {
	if month > 12 {
		return time.Now()
	}
	t := time.Now()
	m := int(t.Month()) + month
	y := t.Year()
	if m > 12 {
		y += 1
		m -= 12
	}
	return time.Date(y, time.Month(m), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)
}

//yyyy-mm-dd 00:00:00:000
func (_dateUtil)  GetStartTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

//yyyy-mm-dd 23:59:59:999
func (_dateUtil) GetEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999, time.Local)
}