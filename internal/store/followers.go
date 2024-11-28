package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type FollowerStore struct {
	db *sql.DB
}

func (s *FollowerStore) Follow(ctx context.Context, followerID int64, userID int64) error {
	query := `
		INSERT INTO followers(user_id, follower_id)
		VALUES ($1, $2)
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)

	if err != nil {
		switch {
		case IsDuplicateKeyError(err):
			return ErrConflict
		default:
			return err
		}
	}

	return nil
}

func (s *FollowerStore) UnFollow(ctx context.Context, followerID int64, userID int64) error {
	query := `
		DELETE FROM followers
		WHERE user_id = $1 AND follower_id = $2
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.ExecContext(
		ctx,
		query,
		userID,
		followerID,
	)

	if err != nil {
		return err
	}

	row, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return ErrNotFound
	}

	return nil
}

// Helper function to check for duplicate key errors
func IsDuplicateKeyError(err error) bool {
	if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
		return true
	}

	return false
}
