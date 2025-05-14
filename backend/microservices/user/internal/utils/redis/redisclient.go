package redis

import (
	"backend/microservices/user/internal/config"
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg config.MyRedis) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Host,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
	}
}

func (c *RedisClient) SetRefreshTokenJTI(ctx context.Context, userID uint64, jti string) error {
	return c.client.Set(ctx, "refresh_token:jti:"+strconv.FormatUint(userID, 10), jti, time.Hour*24*7).Err()
}

func (c *RedisClient) GetRefreshTokenJTI(ctx context.Context, userID uint64) (string, error) {
	return c.client.Get(ctx, "refresh_token:jti:"+strconv.FormatUint(userID, 10)).Result()
}

func (c *RedisClient) DeleteRefreshTokenJTI(ctx context.Context, userID uint64) error {
	return c.client.Del(ctx, "refresh_token:jti:"+strconv.FormatUint(userID, 10)).Err()
}

func (c *RedisClient) SetAccessTokenJTI(ctx context.Context, userID uint64, jti string) error {
	return c.client.Set(ctx, "access_token:jti:"+strconv.FormatUint(userID, 10), jti, time.Minute*15).Err()
}

func (c *RedisClient) GetAccessTokenJTI(ctx context.Context, userID uint64) (string, error) {
	return c.client.Get(ctx, "access_token:jti:"+strconv.FormatUint(userID, 10)).Result()
}

func (c *RedisClient) DeleteAccessTokenJTI(ctx context.Context, userID uint64) error {
	return c.client.Del(ctx, "access_token:jti:"+strconv.FormatUint(userID, 10)).Err()
}
