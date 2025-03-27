package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	cars         = make(map[int]Car)
	users        = make(map[int]User)
	reservations = make(map[int]Reservation)
	mu           = &sync.Mutex{}
	currentUser  = User{}
)

func addCar(make string, model string, year int, license string, rent int) Car {
	id := len(cars) + 1
	cars[id] = Car{id, make, model, year, license, rent, true}
	return cars[id]
}

func addUser(name string, email string, license string) User {
	id := len(users) + 1
	users[id] = User{id, name, email, license}
	return users[id]
}

func (car *Car) reserveCar(userId int, startTime string, endTime string) (Reservation, error) {
	mu.Lock()
	if car.Available {
		car.Available = false
		cars[car.Id] = *car
		id := len(reservations) + 1
		reservations[id] = (Reservation{Id: id, UserId: userId, CarId: car.Id, StartTime: startTime, EndTime: endTime})
		mu.Unlock()
		return reservations[id], nil
	} else {
		mu.Unlock()
		return Reservation{}, errors.New("Car not available")
	}
}

func (reservation *Reservation) cancelReservation() {
	mu.Lock()
	car := cars[reservation.CarId]
	car.Available = true
	delete(reservations, reservation.Id)
	mu.Unlock()
}

func selectUser() User {
	fmt.Println("To use existing user enter 1 else enter 2")
	var selUser int
	fmt.Scanln(&selUser)
	var user User
	if selUser == 1 {
		fmt.Println("Select id from available Users:")
		for user := range users {
			fmt.Printf("name: %v, id: %v\n", users[user].Name, user)
		}
		var id int
		fmt.Scanln(&id)
		user = users[id]
	} else {
		fmt.Println("Creating new User")
		var (
			name    string
			email   string
			license string
		)
		fmt.Println("Enter name: ")
		fmt.Scanln(&name)
		fmt.Println("Enter email: ")
		fmt.Scanln(&email)
		fmt.Println("Enter license: ")
		fmt.Scanln(&license)
		user = addUser(name, email, license)
	}
	return user
}

func startReserving() {
	fmt.Println("Select a car to continue booking")
	for id, car := range cars {
		if car.Available {
			fmt.Println("Enter", id, "to select", car.Make, car.Model, car.Year, car.Rent)
		}
	}
	var selectedCarId int
	fmt.Scanln(&selectedCarId)
	if selectedCar, exists := cars[selectedCarId]; exists {
		reservation, err := selectedCar.reserveCar(currentUser.Id, "20-03-2025", "21-03-2025")
		if err != nil {
			fmt.Println("Error booking", err)
			return
		} else {
			fmt.Println("Booking successful for", selectedCar.Make, selectedCar.Model, "from", reservation.StartTime, "to", reservation.EndTime)
			return
		}
	} else {
		fmt.Println("Invalid input")
		return
	}
}

func showMyBookings() bool {
	if len(reservations) == 0 {
		fmt.Println("No bookings yet")
		return false
	}
	for _, reservation := range reservations {
		if reservation.UserId == currentUser.Id {
			car := cars[reservation.CarId]
			fmt.Println("Booking id: ", reservation.Id, "-", car.Make, car.Model, "from", reservation.StartTime, "to", reservation.EndTime)
		}
	}
	return true
}

func startCancelReservation() {
	fmt.Println("Select booking id from the following")
	has := showMyBookings()
	if !has {
		return
	}
	var reservationId int
	fmt.Scanln(&reservationId)
	if reservation, exists := reservations[reservationId]; exists {
		reservation.cancelReservation()
		fmt.Println("Reservation cancelled")
		return
	} else {
		fmt.Println("Invalid input")
		return
	}
}

func searchCars() {
	var (
		minPrice      int
		inputMaxPrice int
		availability  int
	)
	fmt.Println("Enter min price")
	fmt.Scanln(&minPrice)
	fmt.Println("Enter max price")
	fmt.Scanln(&inputMaxPrice)
	fmt.Println("Enter 1 to show available cars or enter 2 to show unavailable cars or enter 3 to show all cars")
	fmt.Scanln(&availability)

	maxPrice := 1000000
	if inputMaxPrice > 0 {
		maxPrice = inputMaxPrice
	}

	for _, car := range cars {
		if car.Rent >= minPrice && car.Rent <= maxPrice {
			switch availability {
			case 1:
				if car.Available {
					fmt.Println(car.Make, car.Model, car.Year, car.Rent)
				}
			case 2:
				if !car.Available {
					fmt.Println(car.Make, car.Model, car.Year, car.Rent)
				}
			default:
				fmt.Println(car.Make, car.Model, car.Year, car.Rent)
			}
		}

	}
}

func main() {
	addCar("Maruti", "Alto", 2020, "X123", 3000)
	addCar("Tata", "Nano", 2018, "T320", 1000)
	addCar("Mahindra", "Thar", 2023, "M001", 10000)

	addUser("Raju", "raju@ex.com", "ID123")

	currentUser = selectUser()
	for {
		var input int
		fmt.Println("Select what you want to do")
		fmt.Println("Enter 1 to book a car")
		fmt.Println("Enter 2 to cancel a booking")
		fmt.Println("Enter 3 to view my bookings")
		fmt.Println("Enter 4 to search cars")
		fmt.Println("Enter any key to exit")
		fmt.Scanln(&input)
		switch input {
		case 1:
			startReserving()
		case 2:
			startCancelReservation()
		case 3:
			showMyBookings()
		case 4:
			searchCars()
		default:
			return
		}
		fmt.Printf("--------------------------------\n\n")
	}
}
