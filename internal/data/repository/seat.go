package repository

import (
	"context"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatRepository interface {
	GetStatusSeat(cinema_id int, scheduleTime time.Time) (dto.SeatResponse, error)
}

type seatRepository struct {
	DB *pgxpool.Pool
}

func NewSeatRepository(db *pgxpool.Pool) SeatRepository {
	return &seatRepository{DB: db}
}

func (r *seatRepository) GetStatusSeat(cinema_id int, scheduleTime time.Time) (dto.SeatResponse, error) {
	query := `
		SELECT
			c.cinema_id,
			c.name,
			f.film_id,
			f.name,
			c.time,
			c.capacity,
			c.available,
			c.capacity - c.available AS reserved
		FROM films f JOIN cinemas c
		ON f.film_id = c.film_id
		WHERE c.cinema_id = $1 
		AND c.time BETWEEN ($2::timestamp - INTERVAL '1 hour') AND ($2::timestamp + INTERVAL '1 hour')
		ORDER BY ABS(EXTRACT(EPOCH FROM (c.time - $2::timestamp))) ASC
		LIMIT 1
	`

	var seat dto.SeatResponse

	err := r.DB.QueryRow(context.Background(), query, cinema_id, scheduleTime).Scan(&seat.CinemaId, &seat.Name,&seat.FilmId, &seat.Film, &seat.Time, &seat.Capacity, &seat.Available, &seat.Reserved)

	if err != nil {
		return dto.SeatResponse{}, err
	}

	return seat, nil
}
