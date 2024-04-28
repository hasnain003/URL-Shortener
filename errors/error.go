package errors

import "errors"

var (
	ErrInvalidLongUrl     = errors.New("incorrect long url")
	ErrShortUrlExist      = errors.New("short url already exist")
	ErrLongUrlExist       = errors.New("long url already exist")
	ErrInvalidUrl         = errors.New("incorrect url")
	ErrorUrlAlreadyExist  = errors.New("url already exist")
	ErrInvalidShortUrl    = errors.New("incorrect or expired url")
	ErrInvalidRequestBody = errors.New("invalid request body")
)
