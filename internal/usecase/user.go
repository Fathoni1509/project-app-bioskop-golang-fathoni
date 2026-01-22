package usecase

import (
	"errors"
	"project-app-bioskop-golang-fathoni/internal/data/entity"
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/google/uuid"
)

type UserUsecase interface {
	GetUserByToken(token string) (entity.User, error)
	Register(user *dto.UserRegister) error
	Login(user *dto.UserLogin) (string, error)
	Logout(token string) error
}

type userUsecase struct {
	Repo repository.Repository
}

func NewUserUsecase(repo repository.Repository) UserUsecase {
	return &userUsecase{Repo: repo}
}

// usecase get user by token
func (uc *userUsecase) GetUserByToken(token string) (entity.User, error) {
	return uc.Repo.UserRepo.GetUserByToken(token)
}

// usecase register user
func (uc *userUsecase) Register(user *dto.UserRegister) error {
	return uc.Repo.UserRepo.Register(user)
}

// usecase login user
func (uc *userUsecase) Login(user *dto.UserLogin) (string, error) {
	data, err := uc.Repo.UserRepo.GetUser(user)
	if err != nil {
		return "", errors.New("user not found")
	}

	if user.Password != data.Password {
		return "", errors.New("password wrong")
	}

	newToken := uuid.New().String()

	err = uc.Repo.UserRepo.Login(data.UserId, newToken)
	if err != nil {
		return "", errors.New("login failed")
	}

	return newToken, nil
}

// usecase logout user
func (uc *userUsecase) Logout(token string) error {
	return uc.Repo.UserRepo.Logout(token)
}
