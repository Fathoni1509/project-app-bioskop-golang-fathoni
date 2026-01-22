package dto

import "time"

// type BookingRequest struct {
// 	CinemaId  int       `json:"cinema_id" validate:"required,min=0"`
// 	SeatId    int       `json:"seat_id" validate:"required,min=0"`
// 	PaymentId int       `json:"payment_id" validate:"required,min=0"`
// 	CreatedAt time.Time `json:"created_at" validate:"required"`
// 	UserId    int       `json:"user_id" validate:"required,min=0"`
// 	Status    bool      `json:"status" validate:"required"`
// }

type BookingRequest struct {
	CinemaId int `json:"cinema_id" validate:"required,min=0"`
	SeatId   int `json:"seat_id" validate:"required,min=0"`
	// TimeBooking time.Time `json:"time_booking" validate:"required"`
	// DateBooking time.Time `json:"date_booking" validate:"required"`
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

// type BookingResponse struct {
// 	BookingId int       `json:"booking_id"`
// 	CinemaId  int       `json:"cinema_id"`
// 	SeatId    int       `json:"seat_id"`
// 	UserId    int       `json:"user_id"`
// 	PaymentId int       `json:"payment_id"`
// 	Status    bool      `json:"status"`
// 	CreatedAt time.Time `json:"created_at"`
// }
