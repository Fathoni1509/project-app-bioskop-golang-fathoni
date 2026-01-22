package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"project-app-bioskop-golang-fathoni/internal/middleware"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type BookingAdaptor struct {
	BookingUsecase usecase.BookingUsecase
	Config         utils.Configuration
}

func NewBookingAdaptor(bookingUsecase usecase.BookingUsecase, config utils.Configuration) BookingAdaptor {
	return BookingAdaptor{
		BookingUsecase: bookingUsecase,
		Config:         config,
	}
}

// create booking
func (adaptor *BookingAdaptor) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req dto.BookingRequest
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

	// parsing to model booking
	booking := dto.BookingRequest{
		CinemaId:  req.CinemaId,
		SeatId:    req.SeatId,
		PaymentId: req.PaymentId,
		UserId:    userId,
	}

	// create product
	err = adaptor.BookingUsecase.CreateBooking(&booking)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "booking success", nil)
}
