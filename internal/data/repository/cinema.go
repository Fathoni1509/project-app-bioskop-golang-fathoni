package repository

import (
	"context"
	"project-app-bioskop-golang-fathoni/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type CinemaRepository interface {
	GetListCinemas(page, limit int) ([]dto.CinemaResponse, int, error)
	GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error)
}

type cinemaRepository struct {
	DB *pgxpool.Pool
	Logger *zap.Logger
}

func NewCinemaRepository(db *pgxpool.Pool, log *zap.Logger) CinemaRepository {
	return &cinemaRepository{DB: db, Logger: log}
}

func (r *cinemaRepository) GetListCinemas(page, limit int) ([]dto.CinemaResponse, int, error) {
	offset := (page - 1) * limit

	// get total data for pagination
	var total int
	countQuery := `SELECT COUNT(*) FROM films`
	err := r.DB.QueryRow(context.Background(), countQuery).Scan(&total)
	if err != nil {
		r.Logger.Error("error query findall repo ", zap.Error(err))
		return nil, 0, err
	}

	query := `
		SELECT 
			film_id,
			image_url,
			name,
			rating,
			review_count,
			duration_minute,
			genre,
			status
		FROM films
		ORDER BY film_id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(context.Background(), query, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var listCinemas []dto.CinemaResponse
	var list dto.CinemaResponse
	for rows.Next() {
		err := rows.Scan(&list.FilmId, &list.ImageUrl, &list.Film, &list.Rating, &list.ReviewCount, &list.DurationMinute, &list.Genre, &list.Status)
		if err != nil {
			return nil, 0, err
		}
		listCinemas = append(listCinemas, list)
	}

	return listCinemas, total, nil
}

func (r *cinemaRepository) GetListCinemaById(film_id int) (dto.CinemaDetailResponse, error) {
	query := `
		SELECT
			c.cinema_id,
			c.name, 
			f.film_id,
			f.image_url,
			f.name,
			f.rating,
			f.review_count,
			f.duration_minute,
			f.genre,
			f.status,
			c.time,
			f.language,
			f.storyline
		FROM films f JOIN cinemas c
		ON f.film_id = c.film_id
		WHERE f.film_id = $1
	`

	var cinema dto.CinemaDetailResponse

	err := r.DB.QueryRow(context.Background(), query, film_id).Scan(&cinema.CinemaId, &cinema.Name,&cinema.FilmId, &cinema.ImageUrl, &cinema.Film, &cinema.Rating, &cinema.ReviewCount, &cinema.DurationMinute, &cinema.Genre, &cinema.Status, &cinema.Time, &cinema.Language, &cinema.Storyline)

	if err != nil {
		return dto.CinemaDetailResponse{}, err
	}

	return cinema, nil
}
