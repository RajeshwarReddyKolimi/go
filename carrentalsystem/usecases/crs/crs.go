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
	AddCar(car car.Car) (car.Car, error)
	DeleteCar(carId int) error
	ModifyCar(carId int, updatedCar car.Car) (car.Car, error)

	AddUser(user models.User) (models.User, error)
	DeleteUser(userId int) error
	ModifyUser(userId int, updatedUser models.User) (models.User, error)

	MakeReservation(carId int, startTime time.Time, endTime time.Time) (models.Reservation, error)
	CancelReservation(reservationId int) error
	ModifyReservation(reservationId int, reservation models.Reservation) (models.Reservation, error)

	SearchCars(search models.Search) ([]car.Car, error)
	ShowReservations() ([]models.Reservation, error)
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

func (crs *CarRentalSystem) AddCar(newCar car.Car) (car.Car, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	for _, c := range crs.Cars {
		if c.License == newCar.License {
			return car.Car{}, fmt.Errorf("Car license %v already exists", newCar.License)
		}
	}
	crs.lastCarId++
	id := crs.lastCarId
	newCar.Id = id
	crs.Cars[id] = newCar
	return newCar, nil
}

func (crs *CarRentalSystem) DeleteCar(carId int) error {
	if _, ok := crs.Cars[carId]; ok {
		delete(crs.Cars, carId)
		for _, reservation := range crs.Reservations {
			if reservation.CarId == carId {
				delete(crs.Reservations, reservation.Id)
			}
		}
		for userId, user := range crs.Users {
			updatedReservations := []models.Reservation{}
			for _, reservation := range user.Reservations {
				if reservation.CarId != carId {
					updatedReservations = append(updatedReservations, reservation)
				}
			}
			user.Reservations = updatedReservations
			crs.Users[userId] = user
		}

		return nil
	} else {
		return fmt.Errorf("Car with id %v does not exist", carId)
	}
}

func (crs *CarRentalSystem) ModifyCar(carId int, updatedCar car.Car) (car.Car, error) {
	if existingCar, ok := crs.Cars[carId]; ok {
		if &updatedCar.Make != nil {
			existingCar.Make = updatedCar.Make
		}
		if &updatedCar.Model != nil {
			existingCar.Model = updatedCar.Model
		}
		if &updatedCar.Year != nil {
			existingCar.Year = updatedCar.Year
		}
		if &updatedCar.Rent != nil {
			existingCar.Rent = updatedCar.Rent
		}
		crs.Cars[carId] = existingCar
		return existingCar, nil
	}
	return car.Car{}, fmt.Errorf("Car with id %v does not exist", carId)
}

func (crs *CarRentalSystem) AddUser(user models.User) (models.User, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	for _, u := range crs.Users {
		if u.Email == user.Email {
			return models.User{}, fmt.Errorf("User with email %v already exists", user.Email)
		}
	}
	crs.lastUserId++
	id := crs.lastUserId
	user.Id = id
	crs.Users[id] = user
	return user, nil
}

func (crs *CarRentalSystem) DeleteUser(id int) error {
	if _, ok := crs.Users[id]; ok {
		delete(crs.Users, id)
		if crs.CurrentUser.Id == id {
			crs.CurrentUser = models.User{}
		}
		for _, reservation := range crs.Reservations {
			if reservation.UserId == id {
				delete(crs.Reservations, reservation.Id)
			}
		}
		return nil
	} else {
		return fmt.Errorf("User with id %v does not exist", id)
	}
}

func (crs *CarRentalSystem) ModifyUser(userId int, updatedUser models.User) (models.User, error) {
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
		return user, nil
	}
	return models.User{}, fmt.Errorf("User with id %v does not exist", userId)
}

func (crs *CarRentalSystem) MakeReservation(carId int, startTime time.Time, endTime time.Time) (models.Reservation, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	if car, ok := crs.Cars[carId]; ok && car.IsAvailable(startTime, endTime, -1) {
		crs.lastReservationId++
		reservationId := crs.lastReservationId
		reservation := models.Reservation{Id: reservationId, UserId: crs.CurrentUser.Id, CarId: car.Id, StartTime: startTime, EndTime: endTime, Cost: car.Rent * utils.TotalDays(startTime, endTime)}
		crs.Reservations[reservationId] = reservation
		car.Reservations = append(car.Reservations, reservation)
		crs.Cars[carId] = car
		crs.CurrentUser.Reservations = append(crs.CurrentUser.Reservations, reservation)
		crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
		return reservation, nil
	} else {
		return models.Reservation{}, errors.New("Car not available")
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
	for ind, reservation := range car.Reservations {
		if reservation.Id == reservationId {
			car.Reservations = append(car.Reservations[:ind], car.Reservations[ind+1:]...)
			break
		}
	}
	crs.Cars[reservation.CarId] = car
	for ind, reservation := range crs.CurrentUser.Reservations {
		if reservation.Id == reservationId {
			crs.CurrentUser.Reservations = append(crs.CurrentUser.Reservations[:ind], crs.CurrentUser.Reservations[ind+1:]...)
			break
		}
	}
	crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
	delete(crs.Reservations, reservation.Id)
	return nil
}

func (crs *CarRentalSystem) ModifyReservation(reservationId int, reservation models.Reservation) (models.Reservation, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	_, ok := crs.Reservations[reservationId]
	if !ok {
		return models.Reservation{}, fmt.Errorf("Reservation with id %v does not exist", reservationId)
	}
	if reservation.UserId != crs.CurrentUser.Id {
		return models.Reservation{}, fmt.Errorf("Reservation with id %v does not belong to current user", reservationId)
	}
	carId := reservation.CarId
	if carId == 0 {
		carId = crs.Reservations[reservationId].CarId
	}
	car, ok := crs.Cars[carId]
	if !ok {
		return models.Reservation{}, fmt.Errorf("Car with id %v does not exist", carId)
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
		car.Reservations = append(car.Reservations, crs.Reservations[reservationId])
		crs.Cars[carId] = car

		crs.CurrentUser.Reservations = append(crs.CurrentUser.Reservations, crs.Reservations[reservationId])
		crs.Users[crs.CurrentUser.Id] = crs.CurrentUser
		return reservation, nil
	}
	return models.Reservation{}, fmt.Errorf("Car not available")
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

func (crs *CarRentalSystem) ShowReservations() ([]models.Reservation, error) {
	if len(crs.CurrentUser.Reservations) == 0 {
		return []models.Reservation{}, nil
	}
	return crs.CurrentUser.Reservations, nil
}
