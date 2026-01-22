package repository

import (
	"context"
	"errors"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PayRepository interface {
	GetBookingDetail(booking_id int) (dto.BookingDetail, error)
	CreatePay(pay *dto.PayRequest) error
}

type payRepository struct {
	DB *pgxpool.Pool
}

func NewPayRepository(db *pgxpool.Pool) PayRepository {
	return &payRepository{DB: db}
}

// get booking detail
func (r *payRepository) GetBookingDetail(booking_id int) (dto.BookingDetail, error) {
	query := `
		SELECT 
			b.booking_id, 
			c.price,
			b.user_id,
			b.status
		FROM bookings b JOIN cinemas c
		ON b.cinema_id = c.cinema_id
		WHERE booking_id = $1
	`

	var data dto.BookingDetail

	err := r.DB.QueryRow(context.Background(), query, booking_id).Scan(&data.BookingId, &data.TotalPrice, &data.UserId, &data.Status)

	if err != nil {
		return dto.BookingDetail{}, err
	}

	return data, nil
}

// create pay
// change status from false (not yet paid) to true (paid)
func (r *payRepository) CreatePay(pay *dto.PayRequest) error {
	query := `
		UPDATE bookings
		SET payment_id = $1, status = true
		WHERE status = false AND booking_id = $2 AND user_id = $3
	`

	commandTag, err := r.DB.Exec(context.Background(), query,
		pay.PaymentMethodId,
		pay.BookingId,
		pay.UserId,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("booking id not found or booking has been paid")
	}

	return nil
}
