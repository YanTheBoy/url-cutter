package app

import (
	"github.com/iliarkhpv/url-cutter/internal/app/handlers"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) error {
	urlStorage := repository.CreateURLStorage()
	httpHandler := handlers.NewHTTPHandler(urlStorage)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post())
	e.POST("/api/shorten", httpHandler.PostBody())

	e.Logger.Fatal(e.Start(cfg.ServerAddress))

	return nil
}
