package handlers

import (
	"github.com/google/uuid"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h *HTTPHandler) Post(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if len(body) == 0 {
			return c.String(http.StatusBadRequest, "You should set body")
		}
		URL := string(body)
		err = checkURL(URL)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		urlIdentifier := uuid.New().String()
		shortURL := cfg.BaseURL + urlIdentifier
		err = h.urlStorage.Add(urlIdentifier, string(body))
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		return c.String(http.StatusCreated, shortURL)
	}
}
