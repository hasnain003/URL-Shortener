package models

type ShortURLRequest struct {
	LongUrl string `json:"long_url"`
}

type PostResponse struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func CreateResponse(originalUrl, shortUrl string) PostResponse {
	return PostResponse{
		LongUrl:  originalUrl,
		ShortUrl: shortUrl,
	}
}
