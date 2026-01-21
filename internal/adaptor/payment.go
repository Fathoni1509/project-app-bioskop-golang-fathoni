package adaptor

import (
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type PaymentAdaptor struct {
	PaymentUsecase usecase.PaymentUsecase
	Config utils.Configuration
}

func NewPaymentAdaptor(paymentUsecase usecase.PaymentUsecase, config utils.Configuration) PaymentAdaptor {
	return PaymentAdaptor{
		PaymentUsecase: paymentUsecase,
		Config: config,
	}
}

// get list payment methods
func (adaptor *PaymentAdaptor) GetListPaymentMethods(w http.ResponseWriter, r *http.Request) {
	// get data payment methods from service all payments
	payments, err := adaptor.PaymentUsecase.GetListPaymentMethods()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch payment methods: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data payment methods", payments)
}