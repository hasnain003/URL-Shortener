package store

import (
	"context"

	"github.com/URL-Shortener/models"
)

type Store interface {
	FetchOriginalUrl(ctx context.Context, shortUrl string) (string, error)
	FetchShortUrl(ctx context.Context, longUrl string) (string, error)
	InsertShortUrl(ctx context.Context, shortURL, longURL string) error
	IncrementHitCount(ctx context.Context, value string)
	GetTop3(ctx context.Context) []models.MetricsResponse
}
