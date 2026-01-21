package repository

import (
	"context"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository interface {
	CreateBooking(tx pgx.Tx, booking *dto.BookingRequest) error
}

type bookingRepository struct {
	DB *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) BookingRepository {
	return &bookingRepository{DB: db}
}

// create booking
func (r *bookingRepository) CreateBooking(tx pgx.Tx, booking *dto.BookingRequest) error {
	query := `
		INSERT INTO bookings (cinema_id, seat_id, user_id, payment_id, created_at, status)
		VALUES ($1, $2, $3, $4, NOW(), false)
		RETURNING booking_id
	`

	_, err := tx.Exec(context.Background(), query,
		booking.CinemaId,
		booking.SeatId,
		booking.UserId,
		booking.PaymentId,
	)

	if err != nil {
		return err
	}

	return nil
}
