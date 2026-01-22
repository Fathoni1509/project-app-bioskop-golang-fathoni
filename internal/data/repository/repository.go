package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo    UserRepository
	CinemaRepo  CinemaRepository
	SeatRepo    SeatRepository
	PaymentRepo PaymentRepository
	BookingRepo BookingRepository
	HistoryRepo BookingHistoryRepository
	PayRepo     PayRepository
	DB          *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool, log *zap.Logger) Repository {
	return Repository{
		UserRepo:    NewUserRepository(db),
		CinemaRepo:  NewCinemaRepository(db, log),
		SeatRepo:    NewSeatRepository(db),
		PaymentRepo: NewPaymentRepository(db),
		BookingRepo: NewBookingRepository(db),
		HistoryRepo: NewBookingHistoryRepository(db),
		PayRepo:     NewPayRepository(db),
		DB: db,
	}
}
