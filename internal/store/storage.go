package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
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
