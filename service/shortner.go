package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/URL-Shortener/errors"
	"github.com/URL-Shortener/store"
	"github.com/go-kratos/kratos/v2/log"
)

type UrlShortner struct {
	storage store.Store
}

func NewUrlShortner(storage store.Store) *UrlShortner {
	return &UrlShortner{
		storage: storage,
	}
}

func (s *UrlShortner) FetchOriginalUrl(ctx context.Context, shortUrl string) (string, error) {
	originalUrl, err := s.storage.FetchShortUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (s *UrlShortner) CreateShortUrl(ctx context.Context, originalUrl string) (string, error) {
	// checks if the original url is already there in db
	_, err := s.storage.FetchOriginalUrl(ctx, originalUrl)
	if err == nil {
		return "", errors.ErrorUrlAlreadyExist
	} else if err != errors.ErrInvalidShortUrl {
		return "", err
	}
	shortUrl := s.generateUniqueAlias(ctx, originalUrl)
	err = s.storage.InsertShortUrl(ctx, shortUrl, originalUrl)
	if err != nil {
		log.Error("UrlShortner.CreateShortUrl error inserting shortUrl", err)
		return "", err
	}
	return shortUrl, nil
}

// generateUniqueAlias generates a random string of characters for the short URL
func (s *UrlShortner) generateUniqueAlias(ctx context.Context, originalURL string) string {
	var shortURL string
	// Hash the original URL using SHA-256
	hash := sha256.New()
	hash.Write([]byte(originalURL))
	hashValue := hex.EncodeToString(hash.Sum(nil))
	shortURL = hashValue[:7]
	for {
		shortURL += s.generateIcrementalSuffix()
		// if short url already exist then again find the next unique short url
		if _, err := s.storage.FetchShortUrl(ctx, shortURL); err != nil {
			break
		}
	}

	return shortURL
}

func (s *UrlShortner) generateIcrementalSuffix() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var builder strings.Builder
	for i := 0; i < 3; i++ {
		builder.WriteByte(characters[r.Intn(len(characters))])
	}
	return builder.String()
}

func (s *UrlShortner) getDomain(originalUrl string) (string, error) {
	parsedUrl, err := url.Parse(originalUrl)
	if err != nil {
		return "", errors.ErrInvalidLongUrl
	}

	if parsedUrl.Host == "" || parsedUrl.Scheme == "" {
		return "", errors.ErrInvalidLongUrl
	}

	return parsedUrl.Path, nil
}
