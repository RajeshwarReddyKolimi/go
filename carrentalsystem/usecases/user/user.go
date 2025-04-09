package user

import "crs/models"

type UserService interface {
	AddUser(user models.User) error
	DeleteUser(userId int) error
	ModifyUser(userId int, updatedUser models.User) error
	ShowBookings() ([]models.Reservation, error)
}
