package entity

import "time"

type Cinema struct {
	CinemaId  int       `json:"cinema_id"`
	Name      string    `json:"name"`
	FilmId    int       `json:"film_id"`
	Time      time.Time `json:"time"`
	Capacity  int       `json:"capacity"`
	Available int       `json:"available"`
	Price     float32   `json:"price"`
}
