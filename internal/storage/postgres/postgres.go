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
	// TODO implement
	return 0, nil
}

func (s *Storage) GetUser(username, password string) (models.User, error) {
	// TODO implement
	return models.User{
		ID:         0,
		Name:       "Test-name",
		Surname:    "Test-surname",
		Patronymic: "Test-patronymic",
		Phone:      "Test-phone",
		Email:      "Test-email",
		Password:   "Test-password",
		Role:       "Test-role",
		Hospital:   "Test-hospital",
	}, nil
}
