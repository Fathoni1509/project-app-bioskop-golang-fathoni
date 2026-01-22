package dto

import "time"

type BookingRequest struct {
	CinemaId  int       `json:"cinema_id" validate:"required,min=0"`
	SeatId    int       `json:"seat_id" validate:"required,min=0"`
	PaymentId int       `json:"payment_id" validate:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
	UserId    int       `json:"-"`
	Status    bool      `json:"status"`
}

type BookingDetail struct {
	BookingId  int     `json:"booking_id"`
	TotalPrice float32 `json:"total_price"`
	UserId     int     `json:"user_id"`
	Status     bool    `json:"status"`
}
