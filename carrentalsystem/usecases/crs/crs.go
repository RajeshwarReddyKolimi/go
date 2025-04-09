package crs

import (
	"crs/models"
	"crs/usecases/car"
	"crs/utils"
	"errors"
	"fmt"
	"sync"
	"time"
)

type CarRentalSystem struct {
	Cars              map[int]car.Car
	Users             map[int]models.User
	Reservations      map[int]models.Reservation
	lastCarId         int
	lastUserId        int
	lastReservationId int
	mu                *sync.Mutex
	CurrentUser       models.User
}

type ICRS interface {
	AddCar(car car.Car) error
	DeleteCar(carId int) error
	ModifyCar(carId int, updatedCar car.Car) error

	AddUser(user models.User) error
	DeleteUser(userId int) error
	ModifyUser(userId int, updatedUser models.User) error

	MakeReservation(carId int, startTime time.Time, endTime time.Time) error
	CancelReservation(reservationId int) error
	ModifyReservation(reservationId int, reservation models.Reservation) error

	SearchCars(search models.Search) ([]car.Car, error)
	ShowBookings() ([]models.Reservation, error)
}

func New() *CarRentalSystem {
	return &CarRentalSystem{
		Cars:         make(map[int]car.Car),
		Users:        make(map[int]models.User),
		Reservations: make(map[int]models.Reservation),
		mu:           &sync.Mutex{},
		CurrentUser:  models.User{},
	}
}

func (crs *CarRentalSystem) AddCar(car car.Car) error {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	for _, c := range crs.Cars {
		if c.License == car.License {
			return fmt.Errorf("Car license %v already exists", car.License)
		}
	}
	crs.lastCarId++
	id := crs.lastCarId
	car.Id = id
	crs.Cars[id] = car
	return nil
}

func (crs *CarRentalSystem) DeleteCar(carId int) error {
	if _, ok := crs.Cars[carId]; ok {
		delete(crs.Cars, carId)
		return nil
	} else {
		return fmt.Errorf("Car with id %v does not exist", carId)
	}
}

func (crs *CarRentalSystem) ModifyCar(carId int, updatedCar car.Car) error {
	if car, ok := crs.Cars[carId]; ok {
		if &updatedCar.Make != nil {
			car.Make = updatedCar.Make
		}
		if &updatedCar.Model != nil {
			car.Model = updatedCar.Model
		}
		if &updatedCar.Year != nil {
			car.Year = updatedCar.Year
		}
		if &updatedCar.Rent != nil {
			car.Rent = updatedCar.Rent
		}
		crs.Cars[carId] = car
		return nil
	}
	return fmt.Errorf("Car with id %v does not exist", carId)
}

func (crs *CarRentalSystem) AddUser(user models.User) error {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	for _, u := range crs.Users {
		if u.Email == user.Email {
			return fmt.Errorf("User with email %v already exists", user.Email)
		}
	}
	crs.lastUserId++
	id := crs.lastUserId
	user.Id = id
	crs.Users[id] = user
	return nil
}

func (crs *CarRentalSystem) DeleteUser(id int) error {
	if _, ok := crs.Users[id]; ok {
		delete(crs.Users, id)
		return nil
	} else {
		return fmt.Errorf("User with id %v does not exist", id)
	}
}

func (crs *CarRentalSystem) ModifyUser(userId int, updatedUser models.User) error {
	if user, ok := crs.Users[userId]; ok {
		if &updatedUser.Name != nil {
			user.Name = updatedUser.Name
		}
		if &updatedUser.Email != nil {
			user.Email = updatedUser.Email
		}
		if &updatedUser.License != nil {
			user.License = updatedUser.License
		}
		crs.Users[userId] = user
		return nil
	}
	return fmt.Errorf("User with id %v does not exist", userId)
}

func (crs *CarRentalSystem) MakeReservation(carId int, startTime time.Time, endTime time.Time) error {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	if car, ok := crs.Cars[carId]; ok && car.IsAvailable(startTime, endTime, -1) {
		crs.lastReservationId++
		reservationId := crs.lastReservationId
		crs.Reservations[reservationId] = (models.Reservation{Id: reservationId, UserId: crs.CurrentUser.Id, CarId: car.Id, StartTime: startTime, EndTime: endTime, Cost: car.Rent * utils.TotalDays(startTime, endTime)})
		car.Bookings = append(car.Bookings, crs.Reservations[reservationId])
		crs.Cars[carId] = car
		crs.CurrentUser.Bookings = append(crs.CurrentUser.Bookings, crs.Reservations[reservationId])
		crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
		return nil
	} else {
		return errors.New("Car not available")
	}
}

func (crs *CarRentalSystem) CancelReservation(reservationId int) error {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	reservation, ok := crs.Reservations[reservationId]
	if !ok {
		return fmt.Errorf("Reservation with id %v does not exist", reservationId)
	}
	if reservation.UserId != crs.CurrentUser.Id {
		return fmt.Errorf("Reservation with id %v does not belong to current user", reservationId)
	}
	car := crs.Cars[reservation.CarId]
	for ind, booking := range car.Bookings {
		if booking.Id == reservationId {
			car.Bookings = append(car.Bookings[:ind], car.Bookings[ind+1:]...)
			break
		}
	}
	crs.Cars[reservation.CarId] = car
	for ind, booking := range crs.CurrentUser.Bookings {
		if booking.Id == reservationId {
			crs.CurrentUser.Bookings = append(crs.CurrentUser.Bookings[:ind], crs.CurrentUser.Bookings[ind+1:]...)
			break
		}
	}
	crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
	delete(crs.Reservations, reservation.Id)
	return nil
}

func (crs *CarRentalSystem) ModifyReservation(reservationId int, reservation models.Reservation) error {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	_, ok := crs.Reservations[reservationId]
	if !ok {
		return fmt.Errorf("Reservation with id %v does not exist", reservationId)
	}
	if reservation.UserId != crs.CurrentUser.Id {
		return fmt.Errorf("Reservation with id %v does not belong to current user", reservationId)
	}
	carId := reservation.CarId
	if carId == 0 {
		carId = crs.Reservations[reservationId].CarId
	}
	car, ok := crs.Cars[carId]
	if !ok {
		return fmt.Errorf("Car with id %v does not exist", carId)
	}
	startTime := reservation.StartTime
	if startTime.IsZero() {
		startTime = crs.Reservations[reservationId].StartTime
	}
	endTime := reservation.EndTime
	if endTime.IsZero() {
		endTime = crs.Reservations[reservationId].EndTime
	}

	cost := car.Rent * utils.TotalDays(startTime, endTime)
	if car.IsAvailable(startTime, endTime, reservationId) {
		crs.Reservations[reservationId] = models.Reservation{
			Id:        reservationId,
			CarId:     carId,
			StartTime: startTime,
			EndTime:   endTime,
			Cost:      cost,
		}
		car.Bookings = append(car.Bookings, crs.Reservations[reservationId])
		crs.Cars[carId] = car

		crs.CurrentUser.Bookings = append(crs.CurrentUser.Bookings, crs.Reservations[reservationId])
		crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
		return nil
	}
	return fmt.Errorf("Car not available")
}

func (crs *CarRentalSystem) SearchCars(search models.Search) ([]car.Car, error) {
	filteredCars := []car.Car{}
	minPrice := search.MinPrice
	maxPrice := search.MaxPrice
	startTime := search.StartTime
	endTime := search.EndTime
	for _, car := range crs.Cars {
		if (minPrice == 0 || car.Rent >= minPrice) && (maxPrice == 0 || car.Rent <= maxPrice) && (startTime.IsZero() || endTime.IsZero() || car.IsAvailable(startTime, endTime, -1)) {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars, nil
}

func (crs *CarRentalSystem) ShowBookings() ([]models.Reservation, error) {
	if len(crs.CurrentUser.Bookings) == 0 {
		return []models.Reservation{}, nil
	}
	return crs.CurrentUser.Bookings, nil
}
