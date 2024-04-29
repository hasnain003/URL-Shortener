package prommetrics

import (
	"context"
	"net/http"

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

func (m *Metricshandler) GetTop3(ctx *gin.Context) {
	resp := m.urlShortenerService.GetTop3(context.TODO())
	ctx.JSON(http.StatusOK, resp)
}
