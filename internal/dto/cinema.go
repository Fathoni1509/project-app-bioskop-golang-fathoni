package dto

import "time"

type CinemaResponse struct {
	FilmId         int     `json:"film_id"`
	ImageUrl       string  `json:"image_url"`
	Film           string  `json:"film"`
	Rating         float32 `json:"rating"`
	ReviewCount    int     `json:"review_count"`
	DurationMinute int     `json:"duration_minute"`
	Genre          string  `json:"genre"`
	Status         string  `json:"status"`
}

type CinemaDetailResponse struct {
	CinemaId int    `json:"cinema_id"`
	Name     string `json:"name"`
	CinemaResponse
	Time      time.Time `json:"time"`
	Language  string    `json:"language"`
	Storyline string    `json:"storyline"`
}
