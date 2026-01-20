package adaptor

import (
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type Adaptor struct {
	UserAdaptor UserAdaptor
}

func NewAdaptor(usecase usecase.Usecase, config utils.Configuration) Adaptor {
	return Adaptor{
		UserAdaptor: NewUserAdaptor(usecase.UserUsecase, config),
	}
}