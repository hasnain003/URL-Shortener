package store

import (
	"context"
	"errors"
	"testing"

	"github.com/URL-Shortener/models"
)

// MockStorage is a mock implementation of the Store interface for testing purposes
type MockStorage struct{}

func (m *MockStorage) FetchUrl(ctx context.Context, url string) (string, error) {
	// Mocking the behavior of the storage by returning an error if the short URL is "existingUrl"
	if url == "existingUrl" {
		return "https://example.com/existing-url", nil
	} else if url == "https://www.newOriginalUrl.com" {
		return "https://example.com/existing-url", nil
	} else if url == "https://www.existingUrl.com" {
		return "https://example.com/existing-url", nil
	} else if url == "https://www.invalidUrl.com" {
		return "https://example.com/existing-url", nil
	} else if url == "invalidUrl" {
		return "", errors.New("invalid URL")
	}
	// Mocking the behavior of the storage by returning an empty string if the short URL is not found
	return "", nil
}

func (m *MockStorage) InsertShortUrl(ctx context.Context, shortURL, longURL string) error {
	// Mocking the behavior of inserting a short URL
	return nil
}

func (m *MockStorage) IncrementHitCount(ctx context.Context, value string) {
	// Mocking the behavior of incrementing hit count
}

func (m *MockStorage) GetTopK(ctx context.Context, top int) []models.MetricsResponse {
	// Mocking the behavior of getting top K metrics
	return nil
}

func TestFetchOriginalUrl(t *testing.T) {
	// Create an instance of MockStorage
	storage := &MockStorage{}

	// Test case: Fetching an existing short URL
	existingShortUrl := "existingUrl"
	expectedOriginalUrl := "https://example.com/existing-url"
	originalUrl, err := storage.FetchUrl(context.Background(), existingShortUrl)
	if err != nil {
		t.Errorf("Error fetching original URL: %v", err)
	}
	if originalUrl != expectedOriginalUrl {
		t.Errorf("Expected original URL %s, got %s", expectedOriginalUrl, originalUrl)
	}

	// Test case: Fetching a non-existing short URL
	nonExistingShortUrl := "nonExistingUrl"
	_, err = storage.FetchUrl(context.Background(), nonExistingShortUrl)
	if err != nil {
		t.Errorf("Expected no error for non-existing short URL, got %v", err)
	}
}
