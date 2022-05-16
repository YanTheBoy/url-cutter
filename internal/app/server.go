package app

import (
	"github.com/iliarkhpv/url-cutter/internal/app/handlers"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"github.com/labstack/echo/v4"
)

const port = ":8080"

func Run(cfg *config.Config) error {
	urlStorage := repository.CreateURLStorage()
	httpHandler := handlers.NewHTTPHandler(urlStorage)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post())
	e.POST("/api/shorten", httpHandler.PostBody())

	e.Logger.Fatal(e.Start(port))

	return nil
}

/*
Добавьте в сервер новый эндпоинт POST /api/shorten,
принимающий в теле запроса JSON-объект {"url":"<some_url>"}
и возвращающий в ответ объект {"result":"<shorten_url>"}.
Не забудьте добавить тесты на новый эндпоинт, как и на предыдущие.
Помните про HTTP content negotiation, проставляйте правильные
значения в заголовок Content-Type.
*/
