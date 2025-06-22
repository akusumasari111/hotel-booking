package domain

import "time"

type Reservation struct {
	ID        uint      `gorm:"PrimaryKey" json:"id"`
	OrderId   uint      `json:"order_id"`
	RoomId    uint      `json:"room_id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	LessorId  uint      `json:"lessor_id"`
	Price     float64   `json:"price"`
	Qty       uint      `json:"qty"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
