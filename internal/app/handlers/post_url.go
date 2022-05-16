package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Result string `json:"result,omitempty"`
}

type Request struct {
	URL string `json:"url,omitempty"`
}

func (h *HTTPHandler) PostBody() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *Request
		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		URL := host + uuid.New().String()
		err := h.urlStorage.Add(URL, request.URL)
		if err != nil {
			return c.String(http.StatusBadRequest, "error during adding")
		}

		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusCreated, Response{
			Result: URL,
		})
	}
}
