package domain

import "time"

const (
	LESSOR = "lessor"
	TENANT = "tenant"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry"`

	// Relasi
	Carts    []Cart    `json:"carts" gorm:"foreignKey:UserId"`
	Orders   []Order   `json:"orders" gorm:"foreignKey:UserId"`
	Payments []Payment `json:"payments" gorm:"foreignKey:UserId"`

	Verified  bool      `json:"verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:tenant"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
