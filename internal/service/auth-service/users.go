package authservice

import "github.com/vet-clinic-back/sso-service/internal/models"

func (s *AuthService) CreateVet(user models.Vet) (uint, error) {
	return s.storage.CreateVet(user)
}

func (s *AuthService) CreateOwner(user models.Owner) (uint, error) {
	return s.storage.CreateOwner(user)
}

func (s *AuthService) GetOwner(owner models.Owner) (models.Owner, error) {
	return s.storage.GetOwner(owner)
}

func (s *AuthService) GetOwners(filter models.PaginationFilter) ([]models.Owner, error) {
	return s.storage.GetOwners(filter)
}

func (s *AuthService) GetVet(vet models.Vet) (models.Vet, error) {
	return s.storage.GetVet(vet)
}
