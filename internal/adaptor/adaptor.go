package adaptor

import (
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type Adaptor struct {
	UserAdaptor           UserAdaptor
	CinemaAdaptor         CinemaAdaptor
	SeatAdaptor           SeatAdaptor
	PaymentAdaptor        PaymentAdaptor
	BookingAdaptor        BookingAdaptor
	BookingHistoryAdaptor BookingHistoryAdaptor
	PayAdaptor            PayAdaptor
}

func NewAdaptor(usecase usecase.Usecase, config utils.Configuration) Adaptor {
	return Adaptor{
		UserAdaptor:           NewUserAdaptor(usecase.UserUsecase, config),
		CinemaAdaptor:         NewCinemaAdaptor(usecase.CinemaUsecase, config),
		SeatAdaptor:           NewSeatAdaptor(usecase.SeatUsecase, config),
		PaymentAdaptor:        NewPaymentAdaptor(usecase.PaymentUsecase, config),
		BookingAdaptor:        NewBookingAdaptor(usecase.BookingUsecase, config),
		BookingHistoryAdaptor: NewBookingHistoryAdaptor(usecase.BookingHistoryUsecase, config),
		PayAdaptor:            NewPayAdaptor(usecase.PayUsecase, config),
	}
}
