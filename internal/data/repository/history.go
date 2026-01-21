package repository

import (
	"context"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingHistoryRepository interface {
	GetListBookingHistorys(user_id int) ([]dto.HistoryResponse, error)
}

type bookingHistoryRepository struct {
	DB *pgxpool.Pool
}

func NewBookingHistoryRepository(db *pgxpool.Pool) BookingHistoryRepository {
	return &bookingHistoryRepository{DB: db}
}

func (r *bookingHistoryRepository) GetListBookingHistorys(user_id int) ([]dto.HistoryResponse, error) {
	query := `
		SELECT 
			b.booking_id,
			b.cinema_id,
			c.name,
			c.film_id,
			f.name,
			f.duration_minute,
			b.created_at::date,
			c.price,
			f.image_url,
			b.status
		FROM bookings b JOIN cinemas c ON b.cinema_id = c.cinema_id
		JOIN films f ON c.film_id = f.film_id
		WHERE b.user_id = $1
		ORDER BY b.booking_id
	`

	rows, err := r.DB.Query(context.Background(), query, user_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var listBookingHistory []dto.HistoryResponse
	var list dto.HistoryResponse
	for rows.Next() {
		var dateTime time.Time

		err := rows.Scan(&list.BookingId, &list.CinemaId, &list.Name, &list.FilmId, &list.Film, &list.DurationMinute, &dateTime, &list.Price, &list.ImageUrl, &list.Status)
		if err != nil {
			return nil, err
		}

		list.Date = dateTime.Format("2006-01-02")
		
		listBookingHistory = append(listBookingHistory, list)
	}

	return listBookingHistory, nil
}
