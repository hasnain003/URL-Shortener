package shortner

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/URL-Shortener/errors"
	"github.com/URL-Shortener/models"
	"github.com/URL-Shortener/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type UrlShortnerHandler struct {
	urlShortenerService *service.UrlShortner
	domain              string
}

func NewUrlShortnerHandler(service *service.UrlShortner, domain string) *UrlShortnerHandler {
	return &UrlShortnerHandler{
		urlShortenerService: service,
		domain:              domain,
	}
}

func (s *UrlShortnerHandler) Redirect(ctx *gin.Context) {
	shortUrl := ctx.Param("short")
	originalURL, err := s.urlShortenerService.FetchOriginalUrl(shortUrl)
	if err != nil {
		log.Error("ShortnerHandler.Redirect Error fetchong original url", err)
		ctx.JSON(http.StatusNotFound, errors.ErrInvalidShortUrl)
		return
	}
	ctx.JSON(http.StatusMovedPermanently, originalURL)
}

func (s *UrlShortnerHandler) POST(ctx *gin.Context) {
	defer ctx.Request.Body.Close()
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error("ShortnerHandler.Post error while reading request body", err)
		ctx.JSON(http.StatusBadRequest, errors.ErrInvalidRequestBody)
		return
	}

	var req models.ShortURLRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Error("ShortnerHandler.Post error while unmarshalling request body", err)
		ctx.JSON(http.StatusBadRequest, errors.ErrInvalidRequestBody)
		return
	}

	if len(req.LongUrl) == 0 {
		log.Error("ShortnerHandler.Post empty long url", err)
		ctx.JSON(http.StatusBadRequest, errors.ErrInvalidRequestBody)
		return
	}

	shortUrl, err := s.urlShortenerService.CreateShortUrl(req.LongUrl)
	if err != nil {
		log.Error("ShortHandler.Post Error creating short url", err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	resp := models.CreateResponse(req.LongUrl, s.domain+shortUrl)
	ctx.JSON(http.StatusCreated, resp)
}
