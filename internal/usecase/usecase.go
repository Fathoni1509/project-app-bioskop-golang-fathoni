package usecase

import "project-app-bioskop-golang-fathoni/internal/data/repository"

type Usecase struct {
	UserUsecase    UserUsecase
	CinemaUsecase  CinemaUsecase
	SeatUsecase    SeatUsecase
	PaymentUsecase PaymentUsecase
	BookingUsecase BookingUsecase
}

func NewUsecase(repo repository.Repository) Usecase {
	return Usecase{
		UserUsecase:    NewUserUsecase(repo),
		CinemaUsecase:  NewCinemaUsecase(repo),
		SeatUsecase:    NewSeatUsecase(repo),
		PaymentUsecase: NewPaymentUsecase(repo),
		BookingUsecase: NewBookingUsecase(repo),
	}
}
