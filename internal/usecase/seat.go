package usecase

import (
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"time"
)

type SeatUsecase interface {
	GetStatusSeat(cinema_id int, scheduleTime time.Time) (dto.SeatResponse, error)
}

type seatUsecase struct {
	Repo repository.Repository
}

func NewSeatUsecase(repo repository.Repository) SeatUsecase {
	return &seatUsecase{Repo: repo}
}

// usecase get status seat
func (uc *seatUsecase) GetStatusSeat(cinema_id int, scheduleTime time.Time) (dto.SeatResponse, error) {
	return uc.Repo.SeatRepo.GetStatusSeat(cinema_id, scheduleTime)
} 