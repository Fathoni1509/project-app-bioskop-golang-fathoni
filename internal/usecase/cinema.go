package usecase

import (
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
)

type CinemaUsecase interface {
	GetListCinemas() ([]dto.CinemaResponse, error)
	GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error)
}

type cinemaUsecase struct {
	Repo repository.Repository
}

func NewCinemaUsecase(repo repository.Repository) CinemaUsecase {
	return &cinemaUsecase{Repo: repo}
}

// usecase get all cinemas
func (uc *cinemaUsecase) GetListCinemas() ([]dto.CinemaResponse, error) {
	return uc.Repo.CinemaRepo.GetListCinemas()
} 

// usecase get detail cinema
func (uc *cinemaUsecase) GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error) {
	return uc.Repo.CinemaRepo.GetListCinemaById(film_id)
} 