package app

import (
	"github.com/iliarkhpv/url-cutter/internal/app/handlers"
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"github.com/labstack/echo/v4"
)

const port = ":8080"

func Run() error {
	urlStorage := repository.CreateURLStorage()
	httpHandler := handlers.NewHTTPHandler(urlStorage)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post())

	e.Logger.Fatal(e.Start(port))

	return nil
}
