package utils

import (
	"errors"

	"github.com/vet-clinic-back/sso-service/internal/models"
)

var ErrInvalidInputBody = errors.New("invalid input body")

func ValidateSignUpVet(dto models.Vet) error {
	if dto.FullName == "" || dto.Position == "" || dto.ClinicNumber == "" ||
		dto.Email == "" || dto.Password == "" || dto.Phone == "" {
		return ErrInvalidInputBody
	}
	return nil
}

func ValidateSignUpOwner(dto models.Owner) error {
	if dto.FullName == "" || dto.Email == "" || dto.Password == "" || dto.Phone == "" {
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
