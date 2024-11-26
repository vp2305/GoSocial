package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("resource not found")
	ErrConflict = errors.New("resource conflict")

	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
		GetByID(context.Context, int64) (*models.Post, error)
		DeleteByID(context.Context, int64) error
		PatchPostById(context.Context, *models.Post) error
	}
	Comments interface {
		Create(context.Context, *models.Comment) error
		GetByPostID(context.Context, int64) (*[]models.Comment, error)
	}
	Users interface {
		Create(context.Context, *models.User) error
		GetByID(ctx context.Context, userID int64) (*models.User, error)
	}
	Followers interface {
		Follow(context.Context, int64, int64) error
		UnFollow(context.Context, int64, int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db: db},
		Comments:  &CommentStore{db: db},
		Users:     &UserStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}
