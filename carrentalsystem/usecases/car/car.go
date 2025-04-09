package car

import (
	"crs/models"
	"time"
)

type Car struct {
	Id       int
	Make     string
	Model    string
	Year     int
	License  string
	Rent     int
	Bookings []models.Reservation
}

type CarService interface {
	AddCar(car Car) error
	DeleteCar(carId int) error
	ModifyCar(carId int, updatedCar Car) error
	isCarAvailable(startTime time.Time, endTime time.Time) bool
	SearchCars(search models.Search) ([]Car, error)
}

func (car *Car) IsAvailable(startTime time.Time, endTime time.Time, reservationId int) bool {
	for _, booking := range car.Bookings {
		if !(endTime.Before(booking.StartTime) || startTime.After(booking.EndTime)) && booking.Id != reservationId {
			return false
		}
	}
	return true
}
