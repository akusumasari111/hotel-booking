package dto

type CreateCartRequest struct {
	ReservationId uint `json:"room_id"`
	Qty           uint `json:"qty"`
}

type CreatePaymentRequest struct {
	OrderId      string  `json:"order_id"`
	PaymentId    string  `json:"payment_id"`
	ClientSecret string  `json:"client"`
	Amount       float64 `json:"amount"`
	UserId       uint    `json:"user_id"`
}
