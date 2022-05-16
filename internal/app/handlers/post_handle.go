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
		urlIdentifier := uuid.New().String()
		shortURL := host + urlIdentifier
		h.urlStorage.Add(urlIdentifier, string(body))

		return c.String(http.StatusCreated, shortURL)
	}
}

func (h *HTTPHandler) PostBody() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(URL)
		if err := c.Bind(u); err != nil {
			return err
		} else {
			urlIdentifier := uuid.New().String()
			shortURL := host + urlIdentifier
			u.Result = shortURL
			c.Response().Header().Set("Content-Type", "application/json")
			return c.JSON(http.StatusCreated, u)
		}
	}
}
