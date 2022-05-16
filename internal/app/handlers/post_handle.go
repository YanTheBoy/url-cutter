package handlers

import (
	"errors"
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
		urlIdentifier := uuid.New().String()
		shortURL := host + urlIdentifier
		err = h.urlStorage.Add(urlIdentifier, string(body))
		if err != nil {
			return errors.New("error adding")
		}

		return c.String(http.StatusCreated, shortURL)
	}
}
