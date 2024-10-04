package storage

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/storage/postgres"
)

// Iterface to interact with user data
type Auth interface {
	CreateOwner(user models.Owner) (int, error)
	CreateVet(user models.Vet) (int, error)
	GetOwner(username, password string) (models.Owner, error)
	GetVet(username, password string) (models.Vet, error)
}

type Storage struct {
	Auth
}

func New(log *logging.Logger) *Storage {
	return &Storage{
		Auth: postgres.New(log),
	}
}
