package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
)

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByPostID(ctx context.Context, postID int64) (*[]models.Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.username, u.id FROM comments c 
		JOIN users u on u.id = c.user_id
		where c.post_id = $1
		ORDER BY c.created_at DESC;
	`
	rows, err := s.db.QueryContext(
		ctx,
		query,
		postID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}

	for rows.Next() {
		var c models.Comment
		c.User = models.User{}
		err := rows.Scan(
			&c.ID,
			&c.PostID,
			&c.UserID,
			&c.Content,
			&c.CreatedAt,
			&c.User.Username,
			&c.User.ID,
		)

		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return &comments, nil
}
