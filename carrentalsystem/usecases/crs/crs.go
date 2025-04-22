package crs

import (
	"crs/models"
	"crs/types"
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
	lastPaymentId     int
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

	SearchCars(types.Search) ([]car.Car, error)
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

func (crs *CarRentalSystem) getCarByLicense(license string) car.Car {
	for _, car := range crs.Cars {
		if car.License == license {
			return car
		}
	}
	return car.Car{}
}

func (crs *CarRentalSystem) AddCar(newCar car.Car) (car.Car, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	if existingCar := crs.getCarByLicense(newCar.License); existingCar.Id != 0 {
		return car.Car{}, fmt.Errorf("Car with license %v already exists", newCar.License)
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
		return nil
	} else {
		return fmt.Errorf("Car with id %v does not exist", carId)
	}
}

func (crs *CarRentalSystem) ModifyCar(carId int, updatedCar car.Car) (car.Car, error) {
	if existingCar, ok := crs.Cars[carId]; ok {
		if updatedCar.Make != "" {
			existingCar.Make = updatedCar.Make
		}
		if updatedCar.Model != "" {
			existingCar.Model = updatedCar.Model
		}
		if updatedCar.Year != 0 {
			existingCar.Year = updatedCar.Year
		}
		if updatedCar.Rent != 0 {
			existingCar.Rent = updatedCar.Rent
		}
		crs.Cars[carId] = existingCar
		return existingCar, nil
	}
	return car.Car{}, fmt.Errorf("Car with id %v does not exist", carId)
}

func (crs *CarRentalSystem) getUserByEmail(email string) models.User {
	for _, user := range crs.Users {
		if user.Email == email {
			return user
		}
	}
	return models.User{}
}

func (crs *CarRentalSystem) AddUser(user models.User) (models.User, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	if existingUser := crs.getUserByEmail(user.Email); existingUser.Id != 0 {
		return models.User{}, fmt.Errorf("User with email %v already exists", user.Email)
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
		if updatedUser.Name != "" {
			user.Name = updatedUser.Name
		}
		if updatedUser.Email != "" {
			user.Email = updatedUser.Email
		}
		if updatedUser.License != "" {
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
	fmt.Println(crs.CurrentUser)
	if crs.CurrentUser.Id == 0 {
		return models.Reservation{}, fmt.Errorf("No user logged in")
	}
	if car, ok := crs.Cars[carId]; ok && car.IsAvailable(crs.Reservations, startTime, endTime, -1) {
		crs.lastReservationId++
		crs.lastPaymentId++
		reservationId := crs.lastReservationId
		cost := car.Rent * utils.TotalDays(startTime, endTime)
		payment := &models.Payment{
			Id:            crs.lastPaymentId,
			ReservationId: reservationId,
			Amount:        cost,
			Status:        types.Pending,
		}
		reservation := models.Reservation{
			Id:        reservationId,
			UserId:    crs.CurrentUser.Id,
			CarId:     car.Id,
			StartTime: startTime,
			EndTime:   endTime,
			Cost:      cost,
			Payment:   payment,
			Status:    types.Inactive,
		}
		if utils.GetRandomStatus() {
			payment.Status = types.Completed
			reservation.Status = types.Active
			crs.Reservations[reservationId] = reservation
		} else {
			payment.Status = types.Failed
			crs.Reservations[reservationId] = reservation
			return reservation, errors.New("Payment failed")
		}
		return reservation, nil
	} else {
		return models.Reservation{}, errors.New("Car not available")
	}
}

func (crs *CarRentalSystem) cancelPayment(payment *models.Payment) error {
	if payment.Id != 0 && payment.Status == types.Completed {
		payment.Status = types.RefundPending
		if utils.GetRandomStatus() {
			payment.Status = types.RefundCompleted
			return nil
		}
		payment.Status = types.RefundRejected
		return fmt.Errorf("Refund rejected")
	}
	return fmt.Errorf("Invalid input")
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
	if er := crs.cancelPayment(reservation.Payment); er != nil {
		return er
	}
	reservation.Status = types.Inactive
	crs.Reservations[reservationId] = reservation
	return nil
}

func (crs *CarRentalSystem) ModifyReservation(reservationId int, reservation models.Reservation) (models.Reservation, error) {
	crs.mu.Lock()
	defer crs.mu.Unlock()
	existingReservation, ok := crs.Reservations[reservationId]
	if !ok {
		return models.Reservation{}, fmt.Errorf("Reservation with id %v does not exist", reservationId)
	}
	if reservation.UserId != crs.CurrentUser.Id {
		return models.Reservation{}, fmt.Errorf("Reservation with id %v does not belong to current user", reservationId)
	}
	carId := reservation.CarId
	if carId == 0 {
		carId = existingReservation.CarId
	}
	car, ok := crs.Cars[carId]
	if !ok {
		return models.Reservation{}, fmt.Errorf("Car with id %v does not exist", carId)
	}
	startTime := reservation.StartTime
	if startTime.IsZero() {
		startTime = existingReservation.StartTime
	}
	endTime := reservation.EndTime
	if endTime.IsZero() {
		endTime = existingReservation.EndTime
	}

	cost := car.Rent * utils.TotalDays(startTime, endTime)
	if car.IsAvailable(crs.Reservations, startTime, endTime, reservationId) {
		if cost != existingReservation.Cost {
			payment := existingReservation.Payment
			if utils.GetRandomStatus() {
				existingReservation.CarId = carId
				existingReservation.StartTime = startTime
				existingReservation.EndTime = endTime
				existingReservation.Cost = cost
				crs.Reservations[reservationId] = existingReservation
				payment.Status = types.Completed
				crs.Reservations[reservationId] = existingReservation
				return existingReservation, nil
			}
			existingReservation.Status = types.Inactive
			crs.Reservations[reservationId] = existingReservation
			return existingReservation, fmt.Errorf("Cannot modify reservation")
		}
		existingReservation.CarId = carId
		existingReservation.StartTime = startTime
		existingReservation.EndTime = endTime
		existingReservation.Cost = cost
		crs.Reservations[reservationId] = existingReservation
		return existingReservation, nil
	}
	return existingReservation, fmt.Errorf("Car not available")
}

func (crs *CarRentalSystem) SearchCars(minPrice, maxPrice int, startTime, endTime time.Time) ([]car.Car, error) {
	filteredCars := []car.Car{}
	for _, car := range crs.Cars {
		if (minPrice == 0 || car.Rent >= minPrice) && (maxPrice == 0 || car.Rent <= maxPrice) && (startTime.IsZero() || endTime.IsZero() || car.IsAvailable(crs.Reservations, startTime, endTime, -1)) {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars, nil
}

func (crs *CarRentalSystem) ShowReservations() ([]models.Reservation, error) {
	myReservations := []models.Reservation{}
	for _, reservation := range crs.Reservations {
		if reservation.UserId == crs.CurrentUser.Id && reservation.Status == types.Active {
			myReservations = append(myReservations, reservation)
		}
	}
	if len(myReservations) == 0 {
		return []models.Reservation{}, nil
	}
	return myReservations, nil
}
