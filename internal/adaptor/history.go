package adaptor

import (
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/middleware"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
)

type BookingHistoryAdaptor struct {
	BookingHistoryUsecase usecase.BookingHistoryUsecase
	Config utils.Configuration
}

func NewBookingHistoryAdaptor(bookingHistoryUsecase usecase.BookingHistoryUsecase, config utils.Configuration) BookingHistoryAdaptor {
	return BookingHistoryAdaptor{
		BookingHistoryUsecase: bookingHistoryUsecase,
		Config: config,
	}
}

// get list booking by user id use auth
func (adaptor *BookingHistoryAdaptor) GetListBookingHistory(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
        utils.ResponseBadRequest(w, http.StatusUnauthorized, "User ID not found in context", nil)
        return
    }

	// get data booking history from service all booking history
	bookingHistory, err := adaptor.BookingHistoryUsecase.GetListBookingHistory(userId)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch booking history: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data booking history", bookingHistory)
}