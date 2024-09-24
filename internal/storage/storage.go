package storage

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/storage/postgres"
)

// Iterface to interact with user data
type Auth interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Storage struct {
	Auth
}

func New(log *logging.Logger) *Storage {
	return &Storage{
		Auth: postgres.New(log),
	}
}
