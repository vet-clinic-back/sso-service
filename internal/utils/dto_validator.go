package utils

import (
	"errors"

	"github.com/vet-clinic-back/sso-service/internal/models"
)

var ErrInvalidInputBody = errors.New("invalid input body")

func ValidateSignUpDTO(dto models.User) error {
	// TODO implement
	if dto.Name == "" || dto.Surname == "" || dto.Patronymic == "" {
		return ErrInvalidInputBody
	}
	if dto.Email == "" || dto.Password == "" || dto.Phone == "" {
		return ErrInvalidInputBody
	}
	return nil
}

// ValidateSignInDTO - validate sign in dto by email and password
func ValidateSignInDTO(dto models.User) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("invalid input body. email & password required")
	}
	return nil
}
