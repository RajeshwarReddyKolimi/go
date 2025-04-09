package reservation

import (
	"crs/models"
	"time"
)

type ReservationService interface {
	MakeReservation(carId int, startTime, endTime time.Time) (models.Reservation, error)
	CancelReservation(reservationId int) error
	ModifyReservation(reservationId int, carId int, startTime, endTime time.Time) error
	GetReservations() (map[int]models.Reservation, error)
}
