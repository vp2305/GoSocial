package store

import (
	"SocialMedia/internal/models"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/lib/pq"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `
		INSERT INTO users (username, password, email, role_id)
		VALUES ($1, $2, $3, ( SELECT id FROM roles where name = $4 )) 
		RETURNING id, created_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// Default any new user as regular user if not specified
	role := user.Role.Name
	if role == "" {
		role = "user"
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password.Hash,
		user.Email,
		role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Constraint {
			case "users_username_key":
				return ErrDuplicateUsername
			case "users_email_key":
				return ErrDuplicateEmail
			default:
				return err
			}
		}
	}

	return nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *models.User, token string, invitationExp time.Duration) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		if err := s.createUserInvitation(ctx, tx, token, invitationExp, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT users.id, username, email, password, created_at, roles.id, roles.name, roles.level, roles.description FROM users
		JOIN roles ON (users.role_id = roles.id)
		WHERE users.id = $1 AND is_active = true
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user models.User

	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.CreatedAt,
		&user.Role.ID,
		&user.Role.Name,
		&user.Role.Level,
		&user.Role.Description,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users
		WHERE email = $1 AND is_active = true
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user models.User

	err := s.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.Hash,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) Activate(ctx context.Context, token string) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		// 1. find the user that this token belongs to.
		user, err := s.getUserFromInvitations(ctx, tx, token)

		if err != nil {
			return err
		}

		// 2. Update the user
		user.IsActive = true
		if err := s.update(ctx, tx, user); err != nil {
			return err
		}

		// 3. Clean the invitation
		if err := s.deleteUserInvitations(ctx, tx, user.ID); err != nil {
			return err
		}

		return nil
	})
}

func (s *UserStore) Delete(ctx context.Context, userID int64) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}

	}

	return nil
}

func (s *UserStore) getUserFromInvitations(ctx context.Context, tx *sql.Tx, token string) (*models.User, error) {
	query := `
		SELECT u.id, u.username, u.email, u.is_active, u.created_at 
		FROM users u
		JOIN user_invitations ui ON u.id = ui.user_id
		WHERE ui.token = $1 AND ui.expiry > $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := &models.User{}

	// Turn token into hash
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])

	err := tx.QueryRowContext(
		ctx,
		query,
		hashToken,
		time.Now(),
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, exp time.Duration, userID int64) error {
	query := `
		INSERT INTO user_invitations (user_id, token, expiry)
		VALUES ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	//? Need to use tx since this will be running in a transaction
	_, err := tx.ExecContext(
		ctx,
		query,
		userID,
		token,
		time.Now().Add(exp),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) update(ctx context.Context, tx *sql.Tx, user *models.User) error {
	query := `
		UPDATE users 
		SET username = $1, email = $2, is_active = $3 
		WHERE id = $4
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.IsActive,
		user.ID,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *UserStore) deleteUserInvitations(ctx context.Context, tx *sql.Tx, userID int64) error {
	query := `
		DELETE FROM user_invitations
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(
		ctx,
		query,
		userID,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}
