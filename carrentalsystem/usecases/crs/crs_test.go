package crs

import (
	"crs/models"
	"crs/usecases/car"
	"testing"
	"time"
)

func TestAddUser(t *testing.T) {
	testcases := []struct {
		name string
		user models.User
	}{
		{"Valid user", models.User{Name: "John Doe", Email: "john@example.com", License: "XYZ123"}},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			system := New()
			addedUser, err := system.AddUser(test.user)
			if err != nil {
				t.Errorf("Expected nil error, got %v", err)
			}
			if addedUser.Id == 0 {
				t.Error("Expected user ID to be assigned")
			}
		})
	}
}

func TestAddCar(t *testing.T) {
	t.Run("Add valid car", func(t *testing.T) {
		system := New()
		c := car.Car{Make: "Toyota", Model: "Camry", Year: 2022, License: "CAR123", Rent: 100}
		addedCar, err := system.AddCar(c)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if addedCar.Id == 0 {
			t.Error("Expected car ID to be assigned")
		}
	})
}

func TestDeleteCar(t *testing.T) {
	t.Run("Delete added car", func(t *testing.T) {
		system := New()
		c := car.Car{Make: "Nissan", Model: "Altima", Year: 2020, License: "CAR999", Rent: 90}
		addedCar, _ := system.AddCar(c)
		err := system.DeleteCar(addedCar.Id)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
}

func TestModifyCar(t *testing.T) {
	t.Run("Modify added car", func(t *testing.T) {
		system := New()
		c := car.Car{Make: "Ford", Model: "Fiesta", Year: 2019, License: "CAR111", Rent: 80}
		addedCar, _ := system.AddCar(c)
		updatedCar := car.Car{Model: "Focus", Rent: 85}
		modCar, err := system.ModifyCar(addedCar.Id, updatedCar)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if modCar.Model != "Focus" {
			t.Error("Expected updated model to be 'Focus'")
		}
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Delete added user", func(t *testing.T) {
		system := New()
		user := models.User{Name: "Alice", Email: "alice@example.com", License: "ALC123"}
		addedUser, _ := system.AddUser(user)
		err := system.DeleteUser(addedUser.Id)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
	})
}

func TestModifyUser(t *testing.T) {
	t.Run("Modify added user", func(t *testing.T) {
		system := New()
		user := models.User{Name: "Bob", Email: "bob@example.com", License: "BOB123"}
		addedUser, _ := system.AddUser(user)
		updatedUser := models.User{Name: "Bobby", License: "BOB456"}
		modUser, err := system.ModifyUser(addedUser.Id, updatedUser)
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if modUser.Name != "Bobby" {
			t.Error("Expected updated name to be 'Bobby'")
		}
	})
}

func TestSearchCars(t *testing.T) {
	t.Run("Search for cars within price range", func(t *testing.T) {
		system := New()
		c := car.Car{Make: "Audi", Model: "A4", Year: 2020, License: "AUD123", Rent: 130}
		system.AddCar(c)

		cars, err := system.SearchCars(100, 150, time.Time{}, time.Time{})
		if err != nil {
			t.Errorf("Expected nil error, got %v", err)
		}
		if len(cars) == 0 {
			t.Error("Expected to find cars within range")
		}
	})
}
