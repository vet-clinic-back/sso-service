package storage

import (
	"github.com/vet-clinic-back/sso-service/internal/config"
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/storage/postgres"
)

// Iterface to interact with user data
type Auth interface {
	CreateOwner(user models.Owner) (uint, error)
	CreateVet(user models.Vet) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetVet(vet models.Vet) (models.Vet, error)
}

type Pet interface {
	CreatePet(pet models.Pet) (uint, error)
	GetPet(pet models.Pet) (models.Pet, error)
	UpdatePet(pet models.Pet) (models.Pet, error)
	DeletePet(id uint) error
}

type Owner interface {
	CreateOwner(user models.Owner) (uint, error)
	GetOwner(owner models.Owner) (models.Owner, error)
	GetAllOwners() ([]models.Owner, error)
	UpdateOwner(owner models.Owner) (models.Owner, error)
	DeleteOwner(id uint) error
}

type Info interface {
	Owner
	Pet
}

type StorageProcess interface {
	Shutdown() error
}

type Storage struct {
	Auth
	Info
	StorageProcess
}

func New(log *logging.Logger, cfg *config.DbConfig) *Storage {
	pg := postgres.New(log, cfg)
	return &Storage{
		Auth:           pg,
		Info:           pg,
		StorageProcess: pg,
	}
}
