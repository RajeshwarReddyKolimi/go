package types

import "time"

type ReservationStatus int
type PaymentStatus int

const (
	Active ReservationStatus = iota
	Inactive
)

const (
	Pending PaymentStatus = iota
	Completed
	Failed
	RefundPending
	RefundCompleted
	RefundRejected
)

type Search struct {
	MinPrice  int
	MaxPrice  int
	StartTime time.Time
	EndTime   time.Time
}
