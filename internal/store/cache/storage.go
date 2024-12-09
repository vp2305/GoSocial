package cache

import (
	"SocialMedia/internal/models"
	"context"

	"github.com/go-redis/redis/v8"
)

type Storage struct {
	User interface {
		Get(context.Context, int64) (*models.User, error)
		Set(context.Context, *models.User) error
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage{
		User: &UserStore{rdb: rbd},
	}
}
