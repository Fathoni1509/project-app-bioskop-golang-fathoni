package dto

type PaymentRequest struct {
	BookingId       int     `json:"booking_id" validate:"required,min=0"`
	PaymentMethodId int     `json:"payment_method_id" validate:"required,min=0"`
	TotalPrice      float32 `json:"total_price" validate:"required,min=0"`
	Status          bool    `json:"status" validate:"required"`
}
