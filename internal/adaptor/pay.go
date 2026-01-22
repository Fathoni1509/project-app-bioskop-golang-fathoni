package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"project-app-bioskop-golang-fathoni/internal/middleware"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type PayAdaptor struct {
	PayUsecase usecase.PayUsecase
	Config     utils.Configuration
}

func NewPayAdaptor(payUsecase usecase.PayUsecase, config utils.Configuration) PayAdaptor {
	return PayAdaptor{
		PayUsecase: payUsecase,
		Config:     config,
	}
}

// create pay
func (adaptor *PayAdaptor) CreatePay(w http.ResponseWriter, r *http.Request) {
	var req dto.PayRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "error data", nil)
		return
	}

	// validation
	messages, err := utils.ValidateErrors(req)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), messages)
		return
	}

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		utils.ResponseBadRequest(w, http.StatusUnauthorized, "User ID not found in context", nil)
		return
	}

	// parsing to model pay
	pay := dto.PayRequest{
		BookingId:       req.BookingId,
		PaymentMethodId: req.PaymentMethodId,
		TotalPrice:      req.TotalPrice,
		UserId:          userId,
	}

	// create pay
	err = adaptor.PayUsecase.CreatePay(&pay)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "pay success", nil)
}
