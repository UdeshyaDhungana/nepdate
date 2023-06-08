package utils

import (
	"time"
)

var weekDays = []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}
var months = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

func Reduce(arr []int, fn func(int, int) int, initial int) int {
	result := initial
	for _, element := range arr {
		result = fn(result, element)
	}
	return result
}

// Number of days elapsed since that time
func DaysSince(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// Number of days elapsed, including the start date
func DaysSinceInclusive(t time.Time) int {
	return DaysSince(t) + 1
}

func GetWeekDays() []time.Weekday {
	return weekDays
}

func WeekdayShortname(w time.Weekday) string {
	return w.String()[:2]
}

func GetMonths() []int {
	return months
}
