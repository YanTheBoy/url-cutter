package app

import (
	"github.com/iliarkhpv/url-cutter/internal/app/handlers"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/iliarkhpv/url-cutter/internal/middleware"
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) error {
	var storage repository.URLStorage
	var err error

	if cfg.FileStoragePath != "" {
		storage, err = repository.NewInFile(cfg.FileStoragePath)
		if err != nil {
			return err
		}
	} else {
		storage = repository.NewInMemory()
	}
	httpHandler := handlers.NewHTTPHandler(storage)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post(cfg))
	e.POST("/api/shorten", httpHandler.PostBody(cfg))
	e.Use(middleware.Decompress())
	e.Use(middleware.Compress())

	e.Logger.Fatal(e.Start(cfg.ServerAddress))

	return nil
}
