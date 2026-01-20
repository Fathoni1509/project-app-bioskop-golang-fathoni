package usecase

import "project-app-bioskop-golang-fathoni/internal/data/repository"

type Usecase struct {
	UserUsecase   UserUsecase
	CinemaUsecase CinemaUsecase
}

func NewUsecase(repo repository.Repository) Usecase {
	return Usecase{
		UserUsecase:   NewUserUsecase(repo),
		CinemaUsecase: NewCinemaUsecase(repo),
	}
}
