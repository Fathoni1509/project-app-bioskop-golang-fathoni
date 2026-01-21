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
	booking, err := uc.Repo.HistoryRepo.GetListBookingHistorys(pay.BookingId)

	if err != nil {
		return errors.New("booking_id is invalid or does not exist")
	}

	for _, p := range booking {
		if pay.TotalPrice != p.Price {
			return errors.New("price is wrong")
		}
	}

	err = uc.Repo.PayRepo.CreatePay(pay)
	if err != nil {
		return err
	}

	return nil
}
