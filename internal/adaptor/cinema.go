package adaptor

import (
	"net/http"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CinemaAdaptor struct {
	CinemaUsecase usecase.CinemaUsecase
	Config utils.Configuration
}

func NewCinemaAdaptor(cinemaUsecase usecase.CinemaUsecase, config utils.Configuration) CinemaAdaptor {
	return CinemaAdaptor{
		CinemaUsecase: cinemaUsecase,
		Config: config,
	}
}

// get list cinemas
func (adaptor *CinemaAdaptor) GetListCinemas(w http.ResponseWriter, r *http.Request) {
	// get data cinemas from service all cinemas
	cinemas, err := adaptor.CinemaUsecase.GetListCinemas()
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to fetch cinemas: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data cinema", cinemas)
}

// get list cinema by id
func (adaptor *CinemaAdaptor) GetListCinemaById(w http.ResponseWriter, r *http.Request) {
	cinemaIDstr := chi.URLParam(r, "cinemaId")

	cinemaID, err := strconv.Atoi(cinemaIDstr)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid cinema ID", nil)
		return
	}

	response, err := adaptor.CinemaUsecase.GetListCinemaById(cinemaID)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusNotFound, "Cinema not found: "+err.Error(), nil)
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "success get data detail cinema", response)
}