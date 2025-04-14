package utils

import (
	"math/rand"
	"time"
)

const Layout = "2006-01-02"

func ParseStringToDate(date string) time.Time {
	time, _ := time.Parse(Layout, date)
	return time
}

func ParseDateToString(date time.Time) string {
	return date.Format(Layout)
}

func TotalDays(startTime time.Time, endTime time.Time) int {
	return int(endTime.Sub(startTime).Hours()/24) + 1
}

func GetRandomStatus() bool {
	return rand.Int()%2 == 0
}
