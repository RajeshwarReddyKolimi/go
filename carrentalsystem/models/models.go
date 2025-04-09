package models

import (
	"time"
)

type Reservation struct {
	Id        int
	UserId    int
	CarId     int
	StartTime time.Time
	EndTime   time.Time
	Cost      int
}

type User struct {
	Id       int
	Name     string
	Email    string
	License  string
	Bookings []Reservation
}

type Search struct {
	MinPrice  int
	MaxPrice  int
	StartTime time.Time
	EndTime   time.Time
}
