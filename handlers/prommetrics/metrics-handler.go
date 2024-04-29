package prommetrics

import (
	"context"
	"net/http"
	"strconv"

	"github.com/URL-Shortener/service"
	"github.com/gin-gonic/gin"
)

type Metricshandler struct {
	urlShortenerService *service.UrlShortner
}

func NewMetricsHandler(service *service.UrlShortner) *Metricshandler {
	return &Metricshandler{
		urlShortenerService: service,
	}
}

func (m *Metricshandler) GetTopK(ctx *gin.Context) {
	queryParam := ctx.Query("param")
	value, err := strconv.Atoi(queryParam)
	if err != nil {
		// If there's an error parsing the integer, return a bad request response
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid integer value"})
		return
	}

	resp := m.urlShortenerService.GetTopK(context.TODO(), value)
	ctx.JSON(http.StatusOK, resp)
}
