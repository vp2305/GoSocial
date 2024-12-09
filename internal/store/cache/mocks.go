package cache

import (
	"SocialMedia/internal/models"
	"context"
)

func NewMockStore() Storage {
	return Storage{
		User: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Get(ctx context.Context, id int64) (*models.User, error) {
	return nil, nil
}

func (m *MockUserStore) Set(ctx context.Context, user *models.User) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, id int64) {}
