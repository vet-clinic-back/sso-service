package service

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	authservice "github.com/vet-clinic-back/sso-service/internal/service/auth-service"
	infoservice "github.com/vet-clinic-back/sso-service/internal/service/info-service"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

type Auth interface {
	CreateVet(user models.Vet) (uint, error)
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetVet(vet models.Vet) (models.Vet, error)
	//
	CreateToken(id uint, fullname string, isVet bool) (string, error)
	ParseToken(token string) (authservice.Payload, error)
}

type Info interface {
	CreatePet(pet models.Pet) (uint, error)
	GetPet(pet models.Pet) (models.Pet, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DeletePet(id uint) error
	//
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetAllOwners() ([]models.Owner, error)
	UpdateOwner(owner models.Owner) (models.Owner, error)
	DeleteOwner(id uint) error
}

type Service struct {
	Info
	Auth
}

func New(log *logging.Logger, storage storage.Auth, stor storage.Info) *Service {
	return &Service{
		Auth: authservice.New(log, storage),
		Info: infoservice.New(log, stor),
	}
}
