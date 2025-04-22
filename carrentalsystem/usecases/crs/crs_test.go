package crs

import (
	mock_statusgenerator "crs/mocks"
	"crs/models"
	"crs/usecases/car"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func TestAddUser(t *testing.T) {
	t.Run("Add user", func(t *testing.T) {
		crs := New()
		addedUser, err := crs.AddUser(models.User{Name: "John Doe", Email: "john@example.com", License: "XYZ123"})
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if addedUser.Id == 0 {
			t.Error("Expected user ID to be assigned")
		}
	})
}

func TestAddCar(t *testing.T) {
	t.Run("Add car", func(t *testing.T) {
		crs := New()
		c := car.Car{Make: "Toyota", Model: "Camry", Year: 2022, License: "CAR123", Rent: 100}
		addedCar, err := crs.AddCar(c)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if addedCar.Id == 0 {
			t.Error("Expected car ID to be assigned")
		}
	})
}

func TestDeleteCar(t *testing.T) {
	t.Run("Delete car", func(t *testing.T) {
		crs := New()
		c := car.Car{Make: "Nissan", Model: "Altima", Year: 2020, License: "CAR999", Rent: 90}
		addedCar, _ := crs.AddCar(c)
		err := crs.DeleteCar(addedCar.Id)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
}

func TestModifyCar(t *testing.T) {
	t.Run("Modify car", func(t *testing.T) {
		crs := New()
		c := car.Car{Make: "Ford", Model: "Fiesta", Year: 2019, License: "CAR111", Rent: 80}
		addedCar, _ := crs.AddCar(c)
		updatedCar := car.Car{Model: "Focus", Rent: 85}
		modCar, err := crs.ModifyCar(addedCar.Id, updatedCar)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if modCar.Model != "Focus" {
			t.Error("Expected updated model to be 'Focus'")
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Delete user", func(t *testing.T) {
		crs := New()
		user := models.User{Name: "Alice", Email: "alice@example.com", License: "ALC123"}
		addedUser, _ := crs.AddUser(user)
		err := crs.DeleteUser(addedUser.Id)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
}

func TestModifyUser(t *testing.T) {
	t.Run("Modify user", func(t *testing.T) {
		crs := New()
		user := models.User{Name: "Bob", Email: "bob@example.com", License: "BOB123"}
		addedUser, _ := crs.AddUser(user)
		updatedUser := models.User{Name: "Bobby", License: "BOB456"}
		modUser, err := crs.ModifyUser(addedUser.Id, updatedUser)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if modUser.Name != "Bobby" {
			t.Error("Expected updated name to be 'Bobby'")
		}
	})
}

func TestSearchCars(t *testing.T) {
	t.Run("Search cars", func(t *testing.T) {
		crs := New()
		c := car.Car{Make: "Audi", Model: "A4", Year: 2020, License: "AUD123", Rent: 130}
		crs.AddCar(c)
		cars, err := crs.SearchCars(100, 150, time.Time{}, time.Time{})
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if len(cars) == 0 {
			t.Error("Expected to find cars within range")
		}
	})
}

func TestMakeReservation(t *testing.T) {
	t.Run("Make reservation", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		crs := New()
		c := car.Car{Make: "Honda", Model: "Civic", Year: 2022, License: "HON123", Rent: 110}
		addedCar, _ := crs.AddCar(c)
		user := models.User{Name: "Sally", Email: "sally@example.com", License: "SAL123"}
		addedUser, _ := crs.AddUser(user)
		crs.CurrentUser = addedUser
		mockStatus := mock_statusgenerator.NewMockStatusGenerator(controller)
		crs.StatusGenerator = mockStatus
		mockStatus.EXPECT().Generate().Return(true)
		_, err := crs.MakeReservation(addedCar.Id, time.Now(), time.Now().Add(24*time.Hour))
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
}

func TestCancelReservation(t *testing.T) {
	t.Run("Cancel reservation", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		crs := New()
		c := car.Car{Make: "Honda", Model: "Civic", Year: 2022, License: "HON123", Rent: 110}
		addedCar, _ := crs.AddCar(c)
		user := models.User{Name: "Sally", Email: "sally@example.com", License: "SAL123"}
		addedUser, _ := crs.AddUser(user)
		crs.CurrentUser = addedUser
		mockStatus := mock_statusgenerator.NewMockStatusGenerator(controller)
		crs.StatusGenerator = mockStatus
		mockStatus.EXPECT().Generate().Return(true).AnyTimes()
		reservation, err := crs.MakeReservation(addedCar.Id, time.Now(), time.Now().Add(24*time.Hour))
		if err != nil {
			t.Errorf("Expected nil error for making reservation, got %v", err)
		}
		err = crs.CancelReservation(reservation.Id)
		if err != nil {
			t.Errorf("Expected nil error for canceling reservation, got %v", err)
		}
	})
}

func TestModifyReservation(t *testing.T) {
	t.Run("Modify reservation", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		crs := New()
		c := car.Car{Make: "Honda", Model: "Civic", Year: 2022, License: "HON123", Rent: 110}
		addedCar, _ := crs.AddCar(c)
		user := models.User{Name: "Sally", Email: "sally@example.com", License: "SAL123"}
		addedUser, _ := crs.AddUser(user)
		crs.CurrentUser = addedUser
		mockStatus := mock_statusgenerator.NewMockStatusGenerator(controller)
		crs.StatusGenerator = mockStatus
		mockStatus.EXPECT().Generate().Return(true).AnyTimes()
		reservation, err := crs.MakeReservation(addedCar.Id, time.Now(), time.Now().Add(24*time.Hour))
		if err != nil {
			t.Errorf("Expected nil error for making reservation, got %v", err)
		}
		reservation.EndTime = time.Now().Add(48 * time.Hour)
		_, err = crs.ModifyReservation(reservation.Id, reservation)
		if err != nil {
			t.Errorf("Expected nil error for modifying reservation, got %v", err)
		}
	})
}
