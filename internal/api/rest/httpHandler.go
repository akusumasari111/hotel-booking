package rest

import (
	"hotel-booking/config"
	"hotel-booking/internal/helper"
	"hotel-booking/pkg/payment"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RestHandler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helper.Auth
	Config config.AppConfig
	Pc     payment.PaymentClient
}
