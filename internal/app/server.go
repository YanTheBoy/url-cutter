package app

import (
	"github.com/url-cutter/internal/app/httphandlers"
	"github.com/url-cutter/internal/repository"
)

const port = ":8080"

func Run() error {
	urlStorage := repository.CreateURLStorage()
	httpHandler := httphandlers.NewHTTPHandler(urlStorage)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post())

	e.Logger.Fatal(e.Start(port))

	return nil
}