package services

import (
	"github.com/vet-clinic-back/sso-service/internal/models"
	authservice "github.com/vet-clinic-back/sso-service/internal/services/auth-service"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

type Auth interface {
	CreateUser(user models.User) (int, error)
	CreateVeterinarian(vet models.Veterinarian) (int, error)
	CreateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Auth
}

func New(storage storage.Auth) *Service {
	return &Service{
		Auth: authservice.New(storage),
	}
}
