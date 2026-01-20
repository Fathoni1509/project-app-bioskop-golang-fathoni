package usecase

import "project-app-bioskop-golang-fathoni/internal/data/repository"

type Usecase struct {
	UserUsecase UserUsecase
}

func NewUsecase(repo repository.Repository) Usecase {
	return Usecase{
		UserUsecase: NewUserUsecase(repo),
	}
}