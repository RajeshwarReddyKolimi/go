package models

type Car struct {
	Id       int
	Make     string
	Model    string
	Year     int
	License  string
	Rent     int
	Bookings []Reservation
}
