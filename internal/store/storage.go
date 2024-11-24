package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
		GetByID(context.Context, int64) (*models.Post, error)
	}
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db: db},
		Users: &UsersStore{db: db},
	}
}
