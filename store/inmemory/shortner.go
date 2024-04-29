package inmemory

import (
	"context"
	"sort"

	errors "github.com/URL-Shortener/errors"
	"github.com/URL-Shortener/models"
)

type Shortner struct {
	storage map[string]string
	counts  map[string]int
}

func NewShortner() *Shortner {
	return &Shortner{
		storage: make(map[string]string),
		counts:  make(map[string]int),
	}
}

func (s *Shortner) FetchUrl(ctx context.Context, url string) (string, error) {
	if _, ok := s.storage[url]; ok {
		return s.storage[url], nil
	}
	return "", errors.ErrInvalidShortUrl
}

func (s *Shortner) InsertShortUrl(ctx context.Context, shortURL, longURL string) error {
	if shortURL == "" || longURL == "" {
		return errors.ErrInvalidUrl
	}
	if _, ok := s.storage[shortURL]; ok {
		return errors.ErrShortUrlExist
	} else {
		s.storage[shortURL] = longURL
		s.storage[longURL] = shortURL
	}
	return nil
}

func (s *Shortner) IncrementHitCount(ctx context.Context, value string) {
	s.counts[value]++
}

func (s *Shortner) GetTopK(ctx context.Context, top int) []models.MetricsResponse {
	type kv struct {
		Key   string
		Value int
	}
	var sortedValues []kv
	for key, value := range s.counts {
		sortedValues = append(sortedValues, kv{key, value})
	}
	sort.Slice(sortedValues, func(i, j int) bool {
		return sortedValues[i].Value > sortedValues[j].Value
	})

	resp := make([]models.MetricsResponse, 0)
	for i := 0; i < len(sortedValues) && i < top; i++ {
		resp = append(resp, models.NewMetricResponse(sortedValues[i].Value, sortedValues[i].Key))
	}
	return resp
}
