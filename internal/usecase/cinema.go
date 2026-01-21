package usecase

import (
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type CinemaUsecase interface {
	GetListCinemas(page, limit int) (*[]dto.CinemaResponse, *dto.Pagination, error)
	GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error)
}

type cinemaUsecase struct {
	Repo repository.Repository
}

func NewCinemaUsecase(repo repository.Repository) CinemaUsecase {
	return &cinemaUsecase{Repo: repo}
}

// usecase get all cinemas
func (uc *cinemaUsecase) GetListCinemas(page, limit int) (*[]dto.CinemaResponse, *dto.Pagination, error) {
	cinemas, total, err := uc.Repo.CinemaRepo.GetListCinemas(page, limit)

	if err != nil {
		return nil, nil, err
	}

	pagination := dto.Pagination{
		CurrentPage: page,
		Limit: limit,
		TotalPages: utils.TotalPage(limit, int64(total)),
		TotalRecords: total,
	}

	return &cinemas, &pagination, nil
} 

// usecase get detail cinema
func (uc *cinemaUsecase) GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error) {
	return uc.Repo.CinemaRepo.GetListCinemaById(film_id)
} 