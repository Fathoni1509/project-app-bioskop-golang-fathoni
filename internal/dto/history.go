package dto

type HistoryResponse struct {
	BookingId      int     `json:"booking_id"`
	CinemaId       int     `json:"cinema_id"`
	Name           string  `json:"name"`
	FilmId         int     `json:"film_id"`
	Film           string  `json:"film"`
	DurationMinute int     `json:"duration_minute"`
	Date           string  `json:"booking_date"`
	Price          float32 `json:"price"`
	ImageUrl       string  `json:"image_url"`
	Status         bool    `json:"status"`
}
