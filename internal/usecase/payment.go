package usecase

import (
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	"project-app-bioskop-golang-fathoni/internal/dto"
)

type PaymentUsecase interface {
	GetListPaymentMethods() ([]dto.PaymentMethodResponse, error)
}

type paymentUsecase struct {
	Repo repository.Repository
}

func NewPaymentUsecase(repo repository.Repository) PaymentUsecase {
	return &paymentUsecase{Repo: repo}
}

// usecase get all payment methods
func (uc *paymentUsecase) GetListPaymentMethods() ([]dto.PaymentMethodResponse, error) {
	return uc.Repo.PaymentRepo.GetListPaymentMethods()
} 