package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/sso-service/internal/models"
)

const vetsTable = "veterinarian"

func (s *Storage) CreateVet(vet models.Vet) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, phone, password_hash, position, clinic_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", vetsTable)

	var id uint
	err := s.db.QueryRow(query, vet.FullName, vet.Email, vet.Phone, vet.Password, vet.Position, vet.ClinicNumber).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create vet: %w", err)
	}

	return id, nil
}

func (s *Storage) GetVet(vet models.Vet) (models.Vet, error) {
	log := s.log.WithField("op", "Storage.GetVet")

	stmt := s.psql.Select("id", "full_name", "email", "phone").From(vetsTable)

	if vet.ID != 0 {
		stmt = stmt.Where(squirrel.Eq{"id": vet.ID})
	}
	if vet.Email != "" {
		stmt = stmt.Where(squirrel.Eq{"email": vet.Email})
	}
	if vet.Password != "" {
		stmt = stmt.Where(squirrel.Eq{"password_hash": vet.Password})
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Vet{}, err
	}

	log.Debug("query: ", query, " args: ", args)

	err = s.db.QueryRow(query, args...).Scan(&vet.ID, &vet.FullName, &vet.Email, &vet.Phone)
	if err != nil {
		return models.Vet{}, err
	}
	return vet, nil
}
