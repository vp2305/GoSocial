package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postID int64) (*models.Post, error) {
	query := `
		SELECT id, title, user_id, content, tags, created_at, updated_at, version
		FROM posts
		WHERE id = $1
	`

	var post models.Post

	err := s.db.QueryRowContext(
		ctx,
		query,
		postID,
	).Scan(
		&post.ID,
		&post.Title,
		&post.UserID,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostStore) DeleteByID(ctx context.Context, postID int64) error {
	query := `
		DELETE FROM posts 
		WHERE id = $1 
	`
	res, err := s.db.ExecContext(
		ctx,
		query,
		postID,
	)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostStore) PatchPostById(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE posts
		SET title = $2, content = $3, tags = $4, updated_at = NOW(), version = version + 1
		WHERE id = $1 AND version = $5
		RETURNING created_at, updated_at, version
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.ID,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.Version,
	).Scan(
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}
