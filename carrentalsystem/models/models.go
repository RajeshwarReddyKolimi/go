package models

import (
	"crs/types"
	"time"
)

type Reservation struct {
	Id        int
	UserId    int
	CarId     int
	StartTime time.Time
	EndTime   time.Time
	Cost      int
	Payment   *Payment
	Status    types.ReservationStatus
}

type User struct {
	Id      int
	Name    string
	Email   string
	License string
}

type Payment struct {
	Id            int
	ReservationId int
	Amount        int
	Status        types.PaymentStatus
}
