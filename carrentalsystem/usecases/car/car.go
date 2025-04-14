package car

import (
	"crs/models"
	"crs/types"
	"time"
)

type Car struct {
	Id      int
	Make    string
	Model   string
	Year    int
	License string
	Rent    int
}

type CarService interface {
	isCarAvailable(startTime time.Time, endTime time.Time) bool
}

func (car *Car) IsAvailable(reservations map[int]models.Reservation, startTime time.Time, endTime time.Time, reservationId int) bool {
	for _, reservation := range reservations {
		if reservation.CarId == car.Id && reservation.Status == types.Active && !(endTime.Before(reservation.StartTime) || startTime.After(reservation.EndTime)) && reservation.Id != reservationId {
			return false
		}
	}
	return true
}
