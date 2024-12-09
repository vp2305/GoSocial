package store

import (
	"SocialMedia/internal/models"
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct {
}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *models.User) error {
	return nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, user *models.User, token string, invitationExp time.Duration) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (m *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, userID int64) error {
	return nil
}
