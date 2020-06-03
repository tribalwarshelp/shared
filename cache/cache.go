package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	DefaultPrefix = "gqlcache-"
)

type Store interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	FlushAll() error
}

type RediStore struct {
	client redis.UniversalClient
	prefix string
}

func NewRediStore(client redis.UniversalClient) *RediStore {
	return &RediStore{
		client: client,
		prefix: DefaultPrefix,
	}
}

func (s *RediStore) SetPrefix(prefix string) {
	s.prefix = prefix
}

func (s *RediStore) Set(key string, value interface{}, expiration time.Duration) error {
	if err := s.client.Set(context.Background(), s.prefix+key, value, expiration).Err(); err != nil {
		return errors.Wrap(err, "RediStore")
	}
	return nil
}

func (s *RediStore) Get(key string) (string, error) {
	item, err := s.client.Get(context.Background(), s.prefix+key).Result()
	if err != nil {
		return "", errors.Wrap(err, "RediStore")
	}
	return item, nil
}

func (s *RediStore) FlushAll() error {
	keys, err := s.client.Keys(context.Background(), s.prefix+"*").Result()
	if err != nil {
		return errors.Wrap(err, "RediStore")
	}
	if _, err := s.client.Del(context.Background(), keys...).Result(); err != nil {
		return errors.Wrap(err, "RediStore")
	}

	return nil
}
