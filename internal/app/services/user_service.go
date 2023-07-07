package services

import (
	"errors"

	"github.com/go-layer-architecture/internal/domain"
	"github.com/go-layer-architecture/internal/repository"
)

func CreateUser(user *domain.CreateUser) error {
	if err := ValidateUser(user); err != nil {
		return err
	}

	if err := repository.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func ValidateUser(user *domain.CreateUser) error {
	if user.Username == "" {
		return errors.New("name field is required")
	} else if user.Email == "" {
		return errors.New("email field is required")
	} else if user.Password == "" {
		return errors.New("password field is required")
	}

	return nil
}

func FindUsers(user *[]domain.User) error {
	if err := repository.FindUsers(user); err != nil {
		return err
	}

	return nil
}
