package car

import (
	"crs/models"
	"time"
)

type Car struct {
	Id           int
	Make         string
	Model        string
	Year         int
	License      string
	Rent         int
	Reservations []models.Reservation
}

type CarService interface {
	isCarAvailable(startTime time.Time, endTime time.Time) bool
}

func (car *Car) IsAvailable(startTime time.Time, endTime time.Time, reservationId int) bool {
	for _, reservation := range car.Reservations {
		if !(endTime.Before(reservation.StartTime) || startTime.After(reservation.EndTime)) && reservation.Id != reservationId {
			return false
		}
	}
	return true
}
