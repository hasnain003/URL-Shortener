package store

import "context"

type Store interface {
	FetchOriginalUrl(ctx context.Context, shortUrl string) (string, error)
	FetchShortUrl(ctx context.Context, longUrl string) (string, error)
	InsertShortUrl(ctx context.Context, shortURL, longURL string) error
}
