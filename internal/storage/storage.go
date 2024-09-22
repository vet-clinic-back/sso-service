package storage

import "github.com/vet-clinic-back/sso-service/internal/models"

// Iterface to interact with user data
type Auth interface {
	CreateVeterinarian(vet models.Veterinarian) (int, error)
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Storage interface {
	Auth
}
