package main

import "time"

type Reservation struct {
	Id        int
	UserId    int
	CarId     int
	StartTime string
	EndTime   string
}

type Car struct {
	Id        int
	Make      string
	Model     string
	Year      int
	License   string
	Rent      int
	Available bool
}

type User struct {
	Id      int
	Name    string
	Email   string
	License string
}

type Payment struct {
	Status        string
	Time          time.Time
	ReservationId int
}
