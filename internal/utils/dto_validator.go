package utils

import (
	"errors"

	"github.com/vet-clinic-back/sso-service/internal/models"
)

func ValidateSignUpDTO(dto models.User) error {
	// TODO implement
	return nil
}

// ValidateSignInDTO - validate sign in dto by email and password
func ValidateSignInDTO(dto models.User) error {
	if dto.Email == "" || dto.Password == "" {
		return errors.New("invalid input body")
	}
	return nil
}
