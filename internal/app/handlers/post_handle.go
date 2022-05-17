package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type URL struct {
	URL    string `json:"-"`
	Result string `json:"result,omitempty"`
}

func (h *HTTPHandler) Post() echo.HandlerFunc {
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
		shortURL := host + urlIdentifier
		err = h.urlStorage.Add(urlIdentifier, string(body))
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		return c.String(http.StatusCreated, shortURL)
	}
}
