package allservers

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/tribalwarshelp/shared/models"
)

const (
	key = "all_servers"
)

type AllServersCache interface {
	Get() ([]*models.Server, bool)
	Clear() error
}

type allServersCache struct {
	client redis.UniversalClient
}

func New(client redis.UniversalClient) AllServersCache {
	return &allServersCache{client}
}

func (c *allServersCache) Get() ([]*models.Server, bool) {
	sJSON, err := c.client.Get(context.Background(), key).Result()
	if sJSON == "" || err != nil {
		return []*models.Server{}, false
	}

	servers := []*models.Server{}
	if err := json.Unmarshal([]byte(sJSON), &servers); err != nil {
		return []*models.Server{}, false
	}
	return servers, true
}

func (c *allServersCache) Clear() error {
	if err := c.client.Del(context.Background(), key).Err(); err != nil {
		return errors.Wrap(err, "All servers cache")
	}
	return nil
}
