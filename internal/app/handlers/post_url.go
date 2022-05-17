package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	config "github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Response struct {
	Result string `json:"result,omitempty"`
}

type Request struct {
	URL string `json:"url,omitempty"`
}

func (h *HTTPHandler) PostBody(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *Request
		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		err := checkURL(request.URL)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		URLIdentifier := uuid.New().String()
		err = h.urlStorage.Add(URLIdentifier, request.URL)
		if err != nil {
			return c.String(http.StatusBadRequest, "error during adding")
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.JSON(http.StatusCreated, Response{
			Result: strings.Join([]string{cfg.BaseURL, URLIdentifier}, "/"),
		})
	}
}
