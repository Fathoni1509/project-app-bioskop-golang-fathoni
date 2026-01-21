package wire

import (
	// handler
	"project-app-bioskop-golang-fathoni/internal/adaptor"
	"project-app-bioskop-golang-fathoni/internal/data/repository"
	mCostume "project-app-bioskop-golang-fathoni/internal/middleware"
	"project-app-bioskop-golang-fathoni/internal/usecase"
	"project-app-bioskop-golang-fathoni/pkg/utils"
	"sync"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type App struct {
	Route *chi.Mux
	Stop  chan struct{}
	WG    *sync.WaitGroup
}

func Wiring(repo *repository.Repository, log *zap.Logger, config utils.Configuration) *App {
	r := chi.NewRouter()

	// emailJobs := make(chan utils.EmailJob, 10) // BUFFER
	// stop := make(chan struct{})
	// metrics := &utils.Metrics{}
	// wg := &sync.WaitGroup{}

	// wireOrder(r, repo, emailJobs)

	usecase := usecase.NewUsecase(*repo)
	adaptor := adaptor.NewAdaptor(usecase, config)

	mw := mCostume.NewMiddlewareCustome(usecase, log)
	// r.Mount("/api/v1", ApiV1(&handler, mw)
	r.Mount("/api/v1", Apiv1(&adaptor, mw))

	return &App{
		Route: r,
		// Stop:  stop,
		// WG:    wg,
	}
}

func Apiv1(adaptor *adaptor.Adaptor, mw mCostume.MiddlewareCostume) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// r.Use(mw.Logging)

	// auth user
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", adaptor.UserAdaptor.Register)
		r.Post("/login", adaptor.UserAdaptor.Login)
		r.Post("/logout", adaptor.UserAdaptor.Logout)
	})

	// get cinema
	r.Route("/cinemas", func(r chi.Router) {
		r.Get("/", adaptor.CinemaAdaptor.GetListCinemas)
		r.Route("/{cinemaId}", func(r chi.Router) {
			r.Get("/", adaptor.CinemaAdaptor.GetListCinemaById)
		})
	})

	// get seat status
	r.Get("/cinemas/{cinemaId}/seats", adaptor.SeatAdaptor.GetStatusSeat)

	// get payment methods
	r.Get("/payment-methods", adaptor.PaymentAdaptor.GetListPaymentMethods)

	// booking seat
	r.Route("/booking", func(r chi.Router) {
		r.Use(mw.AuthMiddleware)
		r.Post("/", adaptor.BookingAdaptor.CreateBooking)
	})

	// booking history
	r.Route("/user/bookings", func(r chi.Router) {
		r.Use(mw.AuthMiddleware)
		r.Get("/", adaptor.BookingHistoryAdaptor.GetListBookingHistory)
	})

	return r
}
