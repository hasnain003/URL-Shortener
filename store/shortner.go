package store

import errors "github.com/URL-Shortener/errors"

type Shortner struct {
	storage map[string]string
}

func NewShortner() *Shortner {
	return &Shortner{
		storage: make(map[string]string),
	}
}

func (s *Shortner) FetchOriginalUrl(shortUrl string) (string, error) {
	if _, ok := s.storage[shortUrl]; ok {
		return s.storage[shortUrl], nil
	}
	return "", errors.ErrInvalidShortUrl
}

func (s *Shortner) FetchShortUrl(longUrl string) (string, error) {
	if _, ok := s.storage[longUrl]; ok {
		return s.storage[longUrl], nil
	}
	return "", errors.ErrInvalidLongUrl
}

func (s *Shortner) InsertShortUrl(shortURL, longURL string) error {
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
