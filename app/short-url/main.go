package main

import (
	"flag"

	"github.com/URL-Shortener/handlers"
	"github.com/URL-Shortener/handlers/prommetrics"
	"github.com/URL-Shortener/handlers/shortner"
	"github.com/URL-Shortener/service"
	"github.com/URL-Shortener/store/redisstore"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var (
	port          = flag.String("restPort", ":8080", "Port for shortener service")
	domain        = flag.String("domain", "localhost:8080/", "Our application domain name")
	metricsPort   = flag.String("metricsPort", ":8081", "Port for exposing the metrics data")
	redisAddr     = flag.String("redisAddr", "redis:6379", "Redis node address")
	redisPassword = flag.String("redisPassword", "", "Password for redis node")
)

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//inmemory := inmemory.NewShortner()
	//urlService := service.NewUrlShortner(inmemory)
	store := redisstore.NewRedisStore(*redisAddr, *redisPassword)
	urlService := service.NewUrlShortner(store)
	urlHandler := shortner.NewUrlShortnerHandler(urlService, *domain)
	router.GET("/health", handlers.Health)
	router.GET("/:short", urlHandler.Redirect)
	router.POST("v1/create/short-url", urlHandler.POST)
	go router.Run(*port)

	mRouter := gin.Default()
	metricsHandler := prommetrics.NewMetricsHandler(urlService)
	mRouter.GET("v1/metrics/top", metricsHandler.GetTopK) // This api can be make flexible to support multiple top values
	mRouter.Run(*metricsPort)
}
