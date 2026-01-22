package usecase

import (
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
)

type BookingHistoryUsecase interface {
	GetListBookingHistory(user_id int) ([]dto.HistoryResponse, error)
}

type bookingHistoryUsecase struct {
	Repo repository.Repository
}

func NewBookingHistoryUsecase(repo repository.Repository) BookingHistoryUsecase {
	return &bookingHistoryUsecase{Repo: repo}
}

// usecase get all bookingHistory
func (uc *bookingHistoryUsecase) GetListBookingHistory(user_id int) ([]dto.HistoryResponse, error) {
	return uc.Repo.HistoryRepo.GetListBookingHistorys(user_id)
}
