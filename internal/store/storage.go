package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
		GetByID(context.Context, int64) (*models.Post, error)
		DeleteByID(context.Context, int64) error
		PatchPostById(context.Context, *models.Post) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) (*[]models.Comment, error)
	}
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db: db},
		Comments: &CommentStore{db: db},
		Users:    &UserStore{db: db},
	}
}
