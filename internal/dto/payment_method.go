package dto

type PaymentMethodResponse struct {
	PaymentMethodId int    `json:"payment_method_id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
}
