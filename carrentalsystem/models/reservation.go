package models

import "time"

type Reservation struct {
	Id        int
	UserId    int
	CarId     int
	StartTime time.Time
	EndTime   time.Time
	Cost      int
}
