package store

import (
	"context"

	"github.com/URL-Shortener/models"
)

type Store interface {
	FetchUrl(ctx context.Context, url string) (string, error)
	InsertShortUrl(ctx context.Context, shortURL, longURL string) error
	IncrementHitCount(ctx context.Context, value string)
	GetTopK(ctx context.Context, top int) []models.MetricsResponse
}
