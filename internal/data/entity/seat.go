package entity

type Seat struct {
	SeatId   int  `json:"seat_id"`
	CinemaId int  `json:"cinema_id"`
	Status   bool `json:"status"`
}
