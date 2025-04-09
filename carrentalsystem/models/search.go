package models

import "time"

type Search struct {
	MinPrice  int
	MaxPrice  int
	StartTime time.Time
	EndTime   time.Time
}
