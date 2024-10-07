package postgres

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/vet-clinic-back/sso-service/internal/models"
)

const ownersTable = "owner"

func (s *Storage) CreateOwner(owner models.Owner) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, phone, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", ownersTable)

	var id uint
	err := s.db.QueryRow(query, owner.FullName, owner.Email, owner.Phone, owner.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create owner: %w", err)
	}

	return id, nil
}

func (s *Storage) GetOwner(owner models.Owner) (models.Owner, error) {
	log := s.log.WithField("op", "Storage.GetOwner")

	stmt := s.psql.Select("id", "full_name", "email", "phone").From(ownersTable)

	if owner.ID != 0 {
		stmt = stmt.Where(squirrel.Eq{"id": owner.ID})
	}
	if owner.Email != "" {
		stmt = stmt.Where(squirrel.Eq{"email": owner.Email})
	}
	if owner.Password != "" {
		stmt = stmt.Where(squirrel.Eq{"password_hash": owner.Password})
	}

	query, args, err := stmt.ToSql()
	if err != nil {
		return models.Owner{}, err
	}

	log.Debug("query: ", query, " args: ", args)

	err = s.db.QueryRow(query, args...).Scan(&owner.ID, &owner.FullName, &owner.Email, &owner.Phone)
	if err != nil {
		return models.Owner{}, err
	}
	return owner, nil
}
