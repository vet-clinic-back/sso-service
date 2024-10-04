package authservice

import (
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

type AuthService struct {
	log     *logging.Logger
	storage storage.Auth
}

func New(log *logging.Logger, storage storage.Auth) *AuthService {
	return &AuthService{log: log, storage: storage}
}
