package authservice

import "github.com/vet-clinic-back/sso-service/internal/models"

func (s *AuthService) CreateVet(user models.Vet) (int, error) {
	return s.storage.CreateVet(user)
}
func (s *AuthService) CreateOwner(user models.Owner) (int, error) {
	return s.storage.CreateOwner(user)
}
