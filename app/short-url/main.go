package main

import (
	"flag"

	"github.com/URL-Shortener/handlers"
	"github.com/URL-Shortener/handlers/shortner"
	"github.com/URL-Shortener/service"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var (
	port   = flag.String("restPort", ":8080", "Port for shortener service")
	domain = flag.String("domain", "localhost:8080/", "our application domain name")
)

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	urlService := service.NewUrlShortner()
	urlHandler := shortner.NewUrlShortnerHandler(urlService, *domain)
	router.GET("/health", handlers.Health)
	router.GET("/:short", urlHandler.Redirect)
	router.POST("v1/create/short-url", urlHandler.POST)
	router.Run(*port)
}
