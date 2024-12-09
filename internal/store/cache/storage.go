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
		Delete(context.Context, int64)
	}
}

func NewRedisStorage(rbd *redis.Client) Storage {
	return Storage{
		User: &UserStore{rdb: rbd},
	}
}
