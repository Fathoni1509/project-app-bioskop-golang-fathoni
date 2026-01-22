package usecase

import (
	"context"
	"errors"
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
)

type BookingUsecase interface {
	CreateBooking(booking *dto.BookingRequest) error
}

type bookingUsecase struct {
	Repo repository.Repository
}

func NewBookingUsecase(repo repository.Repository) BookingUsecase {
	return &bookingUsecase{Repo: repo}
}

// usecase create booking
func (uc *bookingUsecase) CreateBooking(booking *dto.BookingRequest) error {
	_, err := uc.Repo.CinemaRepo.GetListCinemaById(booking.CinemaId)
	if err != nil {
		return errors.New("cinema_id is invalid or does not exist")
	}

	_, err = uc.Repo.SeatRepo.GetSeat(booking.SeatId, booking.CinemaId)
	if err != nil {
		return errors.New("seat does not exist")
	}

	paymentMethod, err := uc.Repo.PaymentRepo.GetListPaymentMethods()

	searchPayment := false
	for _, p := range paymentMethod {
		if booking.PaymentId == p.PaymentMethodId {
			searchPayment = true
			break
		}
	}

	if searchPayment == false {
		return errors.New("payment_method_id is invalid or does not exist")
	}

	// transaction capacity cinema
	tx, err := uc.Repo.DB.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	// decrease capacity
	err = uc.Repo.SeatRepo.DecreaseCapacity(tx, booking.SeatId)
	if err != nil {
		return err
	}

	err = uc.Repo.BookingRepo.CreateBooking(tx, booking)
	if err != nil {
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}
