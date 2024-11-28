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

	// Duplicate check
	ErrDuplicateUsername = errors.New("a user with that username already exists")
	ErrDuplicateEmail    = errors.New("a user with that email already exists")

	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
		GetByID(context.Context, int64) (*models.Post, error)
		DeleteByID(context.Context, int64) error
		PatchPostById(context.Context, *models.Post) error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) (*[]models.PostWithMetadata, error)
	}
	Comments interface {
		Create(context.Context, *models.Comment) error
		GetByPostID(context.Context, int64) (*[]models.Comment, error)
	}
	Users interface {
		Create(context.Context, *sql.Tx, *models.User) error
		GetByID(context.Context, int64) (*models.User, error)
		CreateAndInvite(context.Context, *models.User, string, time.Duration) error
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

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx_err := tx.Rollback()
		if tx_err != nil {
			return tx_err
		}
		return err
	}

	return tx.Commit()
}
