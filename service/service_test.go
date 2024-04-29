package service

import (
	"context"
	"testing"

	"github.com/URL-Shortener/models"
)

// MockStore is a mock implementation of the Store interface for testing purposes
type MockStore struct{}

func (m *MockStore) FetchUrl(ctx context.Context, url string) (string, error) {
	// Mocking the behavior of the storage by returning an error if the short URL is "existingUrl"
	if url == "existingUrl" {
		return "https://example.com/existing-url", nil
	}
	// Mocking the behavior of the storage by returning an empty string if the URL is not found
	return "", nil
}

func (m *MockStore) InsertShortUrl(ctx context.Context, shortURL, longURL string) error {
	// Mocking the behavior of inserting a short URL
	return nil
}

func (m *MockStore) IncrementHitCount(ctx context.Context, value string) {
	// Mocking the behavior of incrementing hit count
}

func (m *MockStore) GetTopK(ctx context.Context, top int) []models.MetricsResponse {
	// Mocking the behavior of getting top K metrics
	return nil
}

func TestCreateShortUrl(t *testing.T) {
	// Create an instance of UrlShortener with MockStore
	urlShortener := UrlShortner{storage: &MockStore{}}

	// Test case: Creating a short URL for a new original URL
	newOriginalUrl := "newOriginalUrl"
	shortUrl, err := urlShortener.CreateShortUrl(context.Background(), newOriginalUrl)
	if err != nil {
		t.Errorf("Error creating short URL: %v", err)
	}
	if shortUrl == "" {
		t.Error("Expected short URL, got empty string")
	}

	// Test case: Creating a short URL for an existing original URL
	existingOriginalUrl := "existingUrl"
	_, err = urlShortener.CreateShortUrl(context.Background(), existingOriginalUrl)
	if err == nil {
		t.Error("Expected error for existing original URL, but got nil")
	} else if err != nil {
		t.Errorf("Expected error ErrorUrlAlreadyExist, got %v", err)
	}

	// Test case: Creating a short URL for an invalid original URL
	invalidOriginalUrl := "invalidUrl"
	_, err = urlShortener.CreateShortUrl(context.Background(), invalidOriginalUrl)
	if err == nil {
		t.Error("Expected error for invalid original URL, but got nil")
	} else if err.Error() != "UrlShortner.CreateShortUrl error inavlid long url: invalid URL" {
		t.Errorf("Expected error message 'UrlShortner.CreateShortUrl error inavlid long url: invalid URL', got %s", err.Error())
	}
}
