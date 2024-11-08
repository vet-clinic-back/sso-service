package service

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	authservice "github.com/vet-clinic-back/sso-service/internal/service/auth-service"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

type Auth interface {
	CreateVet(user models.Vet) (uint, error)
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetOwners(filter models.PaginationFilter) ([]models.Owner, error)
	GetVet(vet models.Vet) (models.Vet, error)
	//
	CreateToken(id uint, fullname string, isVet bool) (string, error)
	ParseToken(token string) (authservice.Payload, error)
}

type Service struct {
	Auth
}

func New(log *logging.Logger, storage storage.Auth) *Service {
	return &Service{
		Auth: authservice.New(log, storage),
	}
}
