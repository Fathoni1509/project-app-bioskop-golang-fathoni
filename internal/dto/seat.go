package dto

import "time"

type SeatResponse struct {
	CinemaId  int       `json:"cinema_id"`
	Name      string    `json:"name"`
	FilmId    int       `json:"film_id"`
	Film      string    `json:"film"`
	Time      time.Time `json:"time"`
	Capacity  int       `json:"capacity"`
	Available int       `json:"available"`
	Reserved  int       `json:"reserved"`
}
