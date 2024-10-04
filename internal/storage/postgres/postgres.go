package postgres

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/models"
)

type Storage struct {
	log *logging.Logger
	//db *sql.DB
}

func New(log *logging.Logger) *Storage {
	if false {
		log.Fatalf("init postgres failed")
	}
	return &Storage{log: log}
}

func (s *Storage) CreateUser(user models.User) (int, error) {
	op := "Storage.GetUser"
	log := s.log.WithField("op", op)
	// TODO implement
	log.Debug("imitation user creation")
	return 0, nil
}

func (s *Storage) CreateOwner(user models.Owner) (int, error) {
	return 1, nil
}
func (s *Storage) CreateVet(user models.Vet) (int, error) {
	return 1, nil
}
func (s *Storage) GetOwner(username, password string) (models.Owner, error) {

	return models.Owner{
		User: models.User{
			ID:       1,
			FullName: "admin",
			Email:    "example@example.com",
			Phone:    "12345",
			Password: "pass12345",
		},
	}, nil
}
func (s *Storage) GetVet(username, password string) (models.Vet, error) {

	return models.Vet{
		User: models.User{
			ID:       1,
			FullName: "admin",
			Email:    "example@example.com",
			Phone:    "12345",
			Password: "pass12345",
		},
		Position:     "vet",
		ClinicNumber: "qwe1s",
	}, nil
}
