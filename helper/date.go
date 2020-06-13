package helper

import (
	"time"
)

func IsLeapYear(year uint32) bool {
	if year < 1600 {
		return year%4 == 0
	}
	if year%3200 == 0 {
		return false
	}
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	if year%4 == 0 {
		return true
	}

	return false
}

func MonthDays(t time.Time) int {
	y, m, _ := t.Date()
	if m == time.December {
		y++
		m = time.January
	}
	return time.Date(y, m, 0, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Day()
}

