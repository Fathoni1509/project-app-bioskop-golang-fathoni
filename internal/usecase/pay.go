package usecase

import (
	"errors"
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
)

type PayUsecase interface {
	CreatePay(pay *dto.PayRequest) error
}

type payUsecase struct {
	Repo repository.Repository
}

func NewPayUsecase(repo repository.Repository) PayUsecase {
	return &payUsecase{Repo: repo}
}

// usecase create pay
func (uc *payUsecase) CreatePay(pay *dto.PayRequest) error {
	booking, err := uc.Repo.PayRepo.GetBookingDetail(pay.BookingId)

	if err != nil {
		return errors.New("booking not found")
	}

	if booking.UserId != pay.UserId {
		return errors.New("unauthorized: this booking belongs to another user")
	}

	if booking.Status == true {
		return errors.New("booking already paid")
	}

	if float32(booking.TotalPrice) != pay.TotalPrice {
		return errors.New("total price mismatch: please check your payment amount")
	}

	err = uc.Repo.PayRepo.CreatePay(pay)
	if err != nil {
		return err
	}

	return nil
}
