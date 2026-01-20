package entity

import "time"

type Booking struct {
	BookingId int       `json:"booking_id"`
	CinemaId  int       `json:"cinema_id"`
	SeatId    int       `json:"seat_id"`
	UserId    int       `json:"user_id"`
	PaymentId int       `json:"payment_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
