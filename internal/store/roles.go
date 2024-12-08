package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
)

type RoleStore struct {
	db *sql.DB
}

func (r *RoleStore) GetByName(ctx context.Context, roleName string) (*models.Role, error) {
	query := `
	SELECT id, name, level, description 
	FROM roles
	WHERE name = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	role := &models.Role{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		roleName,
	).Scan(
		&role.ID,
		&role.Name,
		&role.Level,
		&role.Description,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return role, nil
}
