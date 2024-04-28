package service

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/URL-Shortener/errors"
	"github.com/URL-Shortener/store/inmemory"
)

type UrlShortner struct {
	storage *inmemory.Shortner
}

func NewUrlShortner() *UrlShortner {
	return &UrlShortner{
		storage: inmemory.NewShortner(),
	}
}

func (s *UrlShortner) FetchOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := s.storage.FetchShortUrl(shortUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}

func (s *UrlShortner) CreateShortUrl(originalUrl string) (string, error) {
	// checks if the original url is already there in db
	if _, err := s.storage.FetchOriginalUrl(originalUrl); err == nil {
		return "", errors.ErrorUrlAlreadyExist
	}
	shortUrl := s.generateUniqueAlias(originalUrl)
	s.storage.InsertShortUrl(shortUrl, originalUrl)

	return shortUrl, nil
}

// generateUniqueAlias generates a random string of characters for the short URL
func (s *UrlShortner) generateUniqueAlias(originalURL string) string {
	var shortURL string
	// Hash the original URL using SHA-256
	hash := sha256.New()
	hash.Write([]byte(originalURL))
	hashValue := hex.EncodeToString(hash.Sum(nil))
	shortURL = hashValue[:7]
	for {
		shortURL += s.generateIcrementalSuffix()
		// if short url already exist then again find the next unique short url
		if _, err := s.storage.FetchShortUrl(shortURL); err != nil {
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
