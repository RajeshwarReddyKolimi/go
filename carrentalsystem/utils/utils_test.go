package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestParseStringToDate(t *testing.T) {
	testcases := []struct {
		input    string
		expected time.Time
		isValid  bool
	}{
		{"2025-04-21", time.Date(2025, 4, 21, 0, 0, 0, 0, time.UTC), true},
		{"2023-01-01", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{"9999-12-31", time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC), true},
		{"0001-01-01", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{"2024-02-29", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), true},
		{"2025-13-01", time.Time{}, false},
		{"abcd-ef-gh", time.Time{}, false},
		{"2025-04-31", time.Time{}, false},
		{"2025-04-00", time.Time{}, false},
		{"2025-04-21 15:04:00", time.Time{}, false},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			got := ParseStringToDate(tc.input)
			fmt.Println("Got:", got)
			if tc.isValid {
				if !got.Equal(tc.expected) {
					t.Errorf("Got %v, want %v", got, tc.expected)
				}
			}
		})
	}
}

func TestParseDateToString(t *testing.T) {
	testcases := []struct {
		input    time.Time
		expected string
	}{
		{time.Date(2025, 4, 21, 15, 30, 0, 0, time.UTC), "2025-04-21"},
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), "2023-01-01"},
		{time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC), "9999-12-31"},
		{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "0001-01-01"},
		{time.Date(2024, 2, 29, 11, 11, 11, 0, time.UTC), "2024-02-29"},
	}

	for _, tc := range testcases {
		t.Run(tc.expected, func(t *testing.T) {
			got := ParseDateToString(tc.input)
			if got != tc.expected {
				t.Errorf("Got %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestTotalDays(t *testing.T) {
	testCases := []struct {
		startTime time.Time
		endTime   time.Time
		want      int
	}{
		{time.Date(2025, 4, 20, 10, 0, 0, 0, time.UTC), time.Date(2025, 4, 22, 10, 0, 0, 0, time.UTC), 3},
		{time.Date(2025, 4, 20, 10, 0, 0, 0, time.UTC), time.Date(2025, 4, 23, 10, 0, 0, 0, time.UTC), 4},
		{time.Date(2025, 4, 30, 10, 0, 0, 0, time.UTC), time.Date(2025, 5, 2, 10, 0, 0, 0, time.UTC), 3},
		{time.Date(2024, 12, 31, 10, 0, 0, 0, time.UTC), time.Date(2025, 1, 4, 10, 0, 0, 0, time.UTC), 5},
	}

	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			got := TotalDays(test.startTime, test.endTime)
			if got != test.want {
				t.Errorf("Got %d, want %d", got, test.want)
			}
		})
	}
}
