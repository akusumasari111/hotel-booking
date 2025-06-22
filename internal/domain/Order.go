package domain

import "time"

type Order struct {
	ID             uint          `gorm:"PrimaryKey" json:"id"`
	UserId         uint          `json:"user_id"`
	Status         string        `json:"status"`
	Amount         float64       `json:"amount"`
	TransactionId  string        `json:"transaction_id"`
	OrderRefNumber string        `json:"order_ref_number"`
	PaymentId      string        `json:"payment_id"`
	Reservations   []Reservation `json:"reservations"`
	CreatedAt      time.Time     `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time     `gorm:"default:current_timestamp"`
}
