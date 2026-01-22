package repository

import (
	"context"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository interface {
	GetListPaymentMethods() ([]dto.PaymentMethodResponse, error)
}

type paymentRepository struct {
	DB *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) PaymentRepository {
	return &paymentRepository{DB: db}
}

// get list payment method
func (r *paymentRepository) GetListPaymentMethods() ([]dto.PaymentMethodResponse, error) {
	query := `
		SELECT 
			payment_id,
			name,
			company
		FROM payment_methods
		ORDER BY payment_id
	`

	rows, err := r.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listPayments []dto.PaymentMethodResponse
	var list dto.PaymentMethodResponse
	for rows.Next() {
		err := rows.Scan(&list.PaymentMethodId, &list.Name, &list.Company)
		if err != nil {
			return nil, err
		}
		listPayments = append(listPayments, list)
	}

	return listPayments, nil
}
