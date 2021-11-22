package ports

import (
	fiber "github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(email string, password string) error
	Register(email string, password string, confirmation string) error
}

type UserRepository interface {
	Login(email string, password string) error
	Register(email string, password string) error
}

type UserHandlers interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}
