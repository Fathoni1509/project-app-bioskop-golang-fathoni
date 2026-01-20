package adaptor

import (
	"fmt"
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type SeatAdaptor struct {
	SeatUsecase usecase.SeatUsecase
	Config      utils.Configuration
}

func NewSeatAdaptor(seatUsecase usecase.SeatUsecase, config utils.Configuration) SeatAdaptor {
	return SeatAdaptor{
		SeatUsecase: seatUsecase,
		Config:      config,
	}
}

// get status seat of cinema
func (adaptor *SeatAdaptor) GetStatusSeat(w http.ResponseWriter, r *http.Request) {
	cinemaIDstr := chi.URLParam(r, "cinemaId")

	cinemaID, err := strconv.Atoi(cinemaIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid seat ID", nil)
		return
	}

	queries := r.URL.Query()

	dateParam := queries.Get("date")
	timeParam := queries.Get("time")

	if timeParam == "" || dateParam == "" {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Missing date or time parameters", nil)
		return
	}

	scheduleParam := fmt.Sprintf("%s %s", dateParam, timeParam)

	layout := "2006-01-02 15:04"
	loc, _ := time.LoadLocation("Asia/Jakarta")

	scheduleTime, err := time.ParseInLocation(layout, scheduleParam, loc)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid date/time format (use YYYY-MM-DD HH:MM)", nil)
		return
	}

	// DEBUG (Bisa dihapus nanti)
    fmt.Printf("Searching CinemaID: %d at Time: %v\n", cinemaID, scheduleTime)

	response, err := adaptor.SeatUsecase.GetStatusSeat(cinemaID, scheduleTime)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Seat not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get status seat", response)
}
