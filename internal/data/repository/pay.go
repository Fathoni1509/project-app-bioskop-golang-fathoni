package repository

import (
	"context"
	"errors"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PayRepository interface {
	CreatePay(pay *dto.PayRequest) error
}

type payRepository struct {
	DB *pgxpool.Pool
}

func NewPayRepository(db *pgxpool.Pool) PayRepository {
	return &payRepository{DB: db}
}

// // get price
// func (r *payRepository) GetPrice(booking_id int) (dto.HistoryResponse, error) {
// 	query := `
// 		SELECT 
// 	`
// }

// create pay
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
