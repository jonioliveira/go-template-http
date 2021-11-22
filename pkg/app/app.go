package app

import (
	"github.com/jonioliveira/go-template/internals/core/services"
	"github.com/jonioliveira/go-template/internals/handlers"
	"github.com/jonioliveira/go-template/internals/repositories"
	"github.com/jonioliveira/go-template/pkg/server"
)

func Start() error {
	mongoConn := "secret🤫"
	// repositories
	userRepository, err := repositories.NewUserRepository(mongoConn)
	if err != nil {
		return err
	}
	// services
	userService := services.NewUserService(userRepository)
	// handlers
	userHandlers := handlers.NewUserHandlers(userService)
	// server
	httpServer := server.NewServer(
		userHandlers,
	)
	httpServer.Initialize()
	return nil
}
