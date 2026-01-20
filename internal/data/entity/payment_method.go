package entity

type PaymentMethod struct {
	PaymentId int    `json:"payment_id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
}
