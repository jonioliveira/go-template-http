package services

import (
	"errors"

	"github.com/jonioliveira/go-template/internals/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.UserService = (*UserService)(nil)

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Login(email string, password string) error {
	err := s.userRepository.Login(email, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Register(email string, password string, confirmation string) error {
	if password != confirmation {
		return errors.New("Password not match")
	}

	err := s.userRepository.Register(email, password)
	if err != nil {
		return nil
	}

	return nil
}
