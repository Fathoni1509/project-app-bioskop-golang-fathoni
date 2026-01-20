package middleware

import (
	"project-app-bioskop-golang-fathoni/internal/usecase"

	"go.uber.org/zap"
)

type MiddlewareCostume struct {
	Usecase usecase.Usecase
	Log     *zap.Logger
}

func NewMiddlewareCustome(usecase usecase.Usecase, log *zap.Logger) MiddlewareCostume {
	return MiddlewareCostume{
		Usecase: usecase,
		Log:     log,
	}
}
