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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

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
			return ErrConflict
		default:
			return err
		}
	}

	return nil
}

func (s *PostStore) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) (*[]models.PostWithMetadata, error) {
	query := `
		SELECT 
			p.id, p.user_id, u.username, p.title, p.content, p.tags, COUNT(c.id) AS comments_count, p.created_at, p.version
		FROM posts p
		LEFT JOIN comments c on c.post_id = p.id
		LEFT JOIN users u on p.user_id = u.id
		JOIN followers f on f.follower_id = p.user_id OR p.user_id = $1
		WHERE f.user_id = $1 OR p.user_id = $1
		GROUP BY p.id, u.username
		ORDER BY p.created_at ` + fq.Sort + `
		LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
		userID,
		fq.Limit,
		fq.Offset,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feed []models.PostWithMetadata
	for rows.Next() {
		var p models.PostWithMetadata

		err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.User.Username,
			&p.Title,
			&p.Content,
			pq.Array(&p.Tags),
			&p.CommentCount,
			&p.CreatedAt,
			&p.Version,
		)

		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return &feed, nil
}
