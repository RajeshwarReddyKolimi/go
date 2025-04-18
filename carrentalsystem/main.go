package main

import (
	"crs/models"
	"crs/types"
	"crs/usecases/car"
	"crs/usecases/crs"
	"crs/utils"
	"log"
	"sync"

	"github.com/fatih/color"
)

func main() {
	color.Green("Welcome to Car Rental System") // Just to demonstrate go mod vendor since there is no other 3rd party dependency
	crs := crs.New()
	if _, err := crs.AddCar(car.Car{Make: "Maruti", Model: "Alto", Year: 2020, License: "X123", Rent: 3000}); err != nil {
		log.Println("Failed to add car:", err)
	} else {
		log.Println("Added car successfully")
	}

	if _, err := crs.AddCar(car.Car{Make: "Hyundai", Model: "i20", Year: 2021, License: "Y456", Rent: 3500}); err != nil {
		log.Println("Failed to add car:", err)
	} else {
		log.Println("Added car successfully")
	}

	if _, err := crs.AddCar(car.Car{Make: "Honda", Model: "Civic", Year: 2019, License: "Z789", Rent: 4500}); err != nil {
		log.Println("Failed to add car:", err)
	} else {
		log.Println("Added car successfully")
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		if _, err := crs.AddCar(car.Car{Make: "Toyota", Model: "Corolla", Year: 2022, License: "T321", Rent: 5000}); err != nil {
			log.Println("Failed to add car:", err)
		} else {
			log.Println("Added car successfully")
		}
		wg.Done()
	}()
	go func() {
		if _, err := crs.AddCar(car.Car{Make: "Toyota", Model: "Carolla", Year: 2022, License: "T321", Rent: 5000}); err != nil {
			log.Println("Failed to add car:", err)
		} else {
			log.Println("Added car successfully")
		}
		wg.Done()
	}()
	wg.Wait()
	if _, err := crs.AddUser(models.User{Name: "Raju", Email: "raju@ex.com", License: "ID123"}); err != nil {
		log.Println("Failed to add user:", err)
	} else {
		log.Println("Added user successfully")
	}
	user, ok := crs.Users[1]
	if !ok {
		log.Println("User doen't exist")
		return
	}
	crs.CurrentUser = user

	if reservation, err := crs.MakeReservation(1, utils.ParseStringToDate("2025-04-10"), utils.ParseStringToDate("2025-04-12")); err != nil {
		log.Println("Failed to make reservation:", err)
	} else if reservation.Payment.Status == types.Failed {
		log.Println("Payment failed")
	} else {
		log.Println("Reservation successful")
	}

	if reservation, err := crs.MakeReservation(2, utils.ParseStringToDate("2025-04-12"), utils.ParseStringToDate("2025-04-13")); err != nil {
		log.Println("Failed to make reservation:", err)
	} else if reservation.Payment.Status == types.Failed {
		log.Println("Payment failed")
	} else {
		log.Println("Reservation successful")
	}

	if reservation, err := crs.MakeReservation(3, utils.ParseStringToDate("2025-04-14"), utils.ParseStringToDate("2025-04-16")); err != nil {
		log.Println("Failed to make reservation:", err)
	} else if reservation.Payment.Status == types.Failed {
		log.Println("Payment failed")
	} else {
		log.Println("Reservation successful")
	}

	if reservation, err := crs.MakeReservation(1, utils.ParseStringToDate("2025-04-11"), utils.ParseStringToDate("2025-04-16")); err != nil {
		log.Println("Failed to make reservation:", err)
	} else if reservation.Payment.Status == types.Failed {
		log.Println("Payment failed")
	} else {
		log.Println("Reservation successful")
	}

	if bookings, err := crs.ShowReservations(); err != nil {
		log.Println("Failed to show bookings:", err)
	} else {
		log.Println("Reservations:")
		for _, booking := range bookings {
			log.Println(crs.Cars[booking.CarId].Make, crs.Cars[booking.CarId].Model, crs.Cars[booking.CarId].Year, crs.Cars[booking.CarId].License, crs.Cars[booking.CarId].Rent, booking.Cost, utils.ParseDateToString(booking.StartTime), utils.ParseDateToString(booking.EndTime))
		}
	}
	if err := crs.CancelReservation(1); err != nil {
		log.Println("Failed to cancel reservation:", err)
	} else {
		log.Println("Reservation cancelled")
	}
	if bookings, err := crs.ShowReservations(); err != nil {
		log.Println("Failed to show bookings:", err)
	} else {
		log.Println("Reservations:")
		for _, booking := range bookings {
			log.Println(crs.Cars[booking.CarId].Make, crs.Cars[booking.CarId].Model, crs.Cars[booking.CarId].Year, crs.Cars[booking.CarId].License, crs.Cars[booking.CarId].Rent, booking.Cost, utils.ParseDateToString(booking.StartTime), utils.ParseDateToString(booking.EndTime))
		}
	}
	if _, err := crs.ModifyReservation(2, models.Reservation{StartTime: utils.ParseStringToDate("2025-04-12"), EndTime: utils.ParseStringToDate("2025-04-15")}); err != nil {
		log.Println("Failed to modify reservation:", err)
	} else {
		log.Println("Reservation modified successfully")
	}

	if bookings, err := crs.ShowReservations(); err != nil {
		log.Println("Failed to show bookings:", err)
	} else {
		log.Println("Reservations:")
		for _, booking := range bookings {
			log.Println(crs.Cars[booking.CarId].Make, crs.Cars[booking.CarId].Model, crs.Cars[booking.CarId].Year, crs.Cars[booking.CarId].License, crs.Cars[booking.CarId].Rent, booking.Cost, utils.ParseDateToString(booking.StartTime), utils.ParseDateToString(booking.EndTime))
		}
	}

	if cars, err := crs.SearchCars(3000, 5000, utils.ParseStringToDate("2025-04-10"), utils.ParseStringToDate("2025-04-12")); err != nil {
		log.Println("Failed to search cars:", err)
	} else {
		log.Println("Filtered cars:")
		for _, car := range cars {
			log.Println(car.Make, car.Model, car.Year, car.License, car.Rent)
		}
	}
}
