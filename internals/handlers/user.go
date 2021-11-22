package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/jonioliveira/go-template/internals/core/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) Login(c *fiber.Ctx) error {
	var email string
	var password string
	// Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandlers) Register(c *fiber.Ctx) error {
	var email string
	var password string
	var confirmPassword string

	// Extract the body and get the email and password
	err := h.userService.Register(email, password, confirmPassword)
	if err != nil {
		return err
	}
	return nil
}
