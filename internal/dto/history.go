package dto

import "time"

type HistoryResponse struct {
	BookingId      int       `json:"booking_id"`
	CinemaId       int       `json:"cinema_id"`
	Name           string    `json:"name"`
	FilmId         int       `json:"film_id"`
	Film           string    `json:"film"`
	DurationMinute int       `json:"duration_minute"`
	Date           time.Time `json:"time"`
	ImageUrl       string    `json:"image_url"`
	Status         bool      `json:"status"`
}
