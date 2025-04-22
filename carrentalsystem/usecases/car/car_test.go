package car

import (
	"crs/models"
	"crs/types"
	"testing"
	"time"
)

func TestCar_IsAvailable(t *testing.T) {
	car := &Car{
		Id:      1,
		Make:    "Toyota",
		Model:   "Corolla",
		Year:    2020,
		License: "ABC123",
		Rent:    100,
	}

	base := time.Date(2025, 4, 20, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name          string
		reservations  map[int]models.Reservation
		startTime     time.Time
		endTime       time.Time
		reservationId int
		want          bool
	}{
		{
			name: "no conflicts",
			reservations: map[int]models.Reservation{
				1: {
					Id:        1,
					CarId:     1,
					Status:    types.Active,
					StartTime: base.Add(-3 * time.Hour),
					EndTime:   base.Add(-2 * time.Hour),
				},
			},
			startTime:     base,
			endTime:       base.Add(1 * time.Hour),
			reservationId: 0,
			want:          true,
		},
		{
			name: "conflicting reservation",
			reservations: map[int]models.Reservation{
				2: {
					Id:        2,
					CarId:     1,
					Status:    types.Active,
					StartTime: base,
					EndTime:   base.Add(2 * time.Hour),
				},
			},
			startTime:     base.Add(1 * time.Hour),
			endTime:       base.Add(3 * time.Hour),
			reservationId: 0,
			want:          false,
		},
		{
			name: "conflicting reservation but same ID (should ignore)",
			reservations: map[int]models.Reservation{
				3: {
					Id:        42,
					CarId:     1,
					Status:    types.Active,
					StartTime: base,
					EndTime:   base.Add(2 * time.Hour),
				},
			},
			startTime:     base.Add(1 * time.Hour),
			endTime:       base.Add(3 * time.Hour),
			reservationId: 42,
			want:          true,
		},
		{
			name: "inactive reservation",
			reservations: map[int]models.Reservation{
				4: {
					Id:        4,
					CarId:     1,
					Status:    types.Inactive,
					StartTime: base,
					EndTime:   base.Add(2 * time.Hour),
				},
			},
			startTime:     base.Add(1 * time.Hour),
			endTime:       base.Add(3 * time.Hour),
			reservationId: 0,
			want:          true,
		},
		{
			name: "reservation for different car",
			reservations: map[int]models.Reservation{
				5: {
					Id:        5,
					CarId:     2,
					Status:    types.Active,
					StartTime: base,
					EndTime:   base.Add(2 * time.Hour),
				},
			},
			startTime:     base.Add(1 * time.Hour),
			endTime:       base.Add(3 * time.Hour),
			reservationId: 0,
			want:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := car.IsAvailable(test.reservations, test.startTime, test.endTime, test.reservationId)
			if got != test.want {
				t.Errorf("IsAvailable() = %v, want %v", got, test.want)
			}
		})
	}
}
