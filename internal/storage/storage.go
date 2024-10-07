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

type StorageProcess interface {
	Shutdown() error
}

type Storage struct {
	Auth
	StorageProcess
}

func New(log *logging.Logger, cfg *config.DbConfig) *Storage {
	pg := postgres.New(log, cfg)
	return &Storage{
		Auth:           pg,
		StorageProcess: pg,
	}
}
