package data

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

type JwtCache struct {
	client rueidis.Client
}

func NewJwtCache(client rueidis.Client) *JwtCache {
	return &JwtCache{client: client}
}

func (j *JwtCache) Get(ctx context.Context, key string) (string, error) {
	return j.client.Do(ctx, j.client.B().Get().Key(key).Build()).ToString()
}

func (j *JwtCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return j.client.Do(ctx, j.client.B().Set().Key(key).Value(value).Ex(expiration).Build()).Error()
}

func (j *JwtCache) Del(ctx context.Context, key string) error {
	return j.client.Do(ctx, j.client.B().Del().Key(key).Build()).Error()
}
