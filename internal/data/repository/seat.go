package repository

import (
	"context"
	"errors"
	"project-app-bioskop-golang-fathoni/internal/data/entity"
	"project-app-bioskop-golang-fathoni/internal/dto"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SeatRepository interface {
	GetStatusSeat(cinema_id int, scheduleTime time.Time) (dto.SeatResponse, error)
	GetSeat(seat_id, cinema_id int) (entity.Seat, error)
	DecreaseCapacity(tx pgx.Tx, seat_id int) error
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
			COUNT(*) AS capacity,
			--c.capacity,
			SUM((NOT s.status)::int) AS available,
			--c.available,
			SUM(s.status::int) AS reserved
			--c.capacity - c.available AS reserved
		FROM films f JOIN cinemas c
		ON f.film_id = c.film_id
		JOIN seats s
		ON c.cinema_id = s.cinema_id
		WHERE c.cinema_id = $1 
		AND c.time BETWEEN ($2::timestamp - INTERVAL '1 hour') AND ($2::timestamp + INTERVAL '1 hour')
		GROUP BY c.cinema_id, c.name, f.film_id, f.name, c.time
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

func (r *seatRepository) GetSeat(seat_id, cinema_id int) (entity.Seat, error) {
	query := `
		SELECT
			seat_id,
			cinema_id
		FROM seats
		WHERE seat_id = $1 AND cinema_id = $2 
	`

	var seat entity.Seat

	err := r.DB.QueryRow(context.Background(), query, seat_id, cinema_id).Scan(&seat.SeatId, &seat.CinemaId)

	if err != nil {
		return entity.Seat{}, err
	}

	return seat, nil
}

func (r *seatRepository) DecreaseCapacity(tx pgx.Tx, seat_id int) error {
	query := `UPDATE seats
	SET status = true
	WHERE seat_id = $1 AND status = false`

	commandTag, err := tx.Exec(context.Background(), query, seat_id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
        return errors.New("seat not available")
    }

	return nil
}