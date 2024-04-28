package redisstore

import (
	"context"

	"github.com/URL-Shortener/errors"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(address, password string) *RedisStore {
	return &RedisStore{
		client: GetClient(address, password),
	}
}

func (r *RedisStore) FetchOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	resp, err := r.client.Get(ctx, shortUrl).Result()
	if err != nil && err != redis.Nil {
		log.Error("RedisStore.fetchoriginalUrl error in Get call", err)
		return resp, err
	} else if err == redis.Nil {
		return "", errors.ErrInvalidShortUrl
	}
	return resp, err
}

func (r *RedisStore) FetchShortUrl(ctx context.Context, longUrl string) (string, error) {
	resp, err := r.client.Get(ctx, longUrl).Result()
	if err != nil && err != redis.Nil {
		log.Error("RedisStore.fetchoriginalUrl error in Get call", err)
		return resp, err
	} else if err == redis.Nil {
		return "", errors.ErrInvalidShortUrl
	}
	return resp, err
}

func (r *RedisStore) InsertShortUrl(ctx context.Context, shortURL, longURL string) error {
	if shortURL == "" || longURL == "" {
		return errors.ErrInvalidUrl
	}
	if _, err := r.FetchShortUrl(ctx, shortURL); err == nil {
		return errors.ErrShortUrlExist
	} else {
		err := r.client.Set(ctx, shortURL, longURL, 0).Err()
		if err != nil {
			log.Error("RedisStore.InsertShortUrl Error Inserting short url", err)
			return err
		}
		err = r.client.Set(ctx, longURL, shortURL, 0).Err()
		if err != nil {
			log.Error("RedisStore.InsertShortUrl Error Inserting long url", err)
			return err
		}
	}
	return nil
}
