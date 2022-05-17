package handlers

import (
	"bytes"
	"github.com/iliarkhpv/url-cutter/internal/cfg"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShorten(t *testing.T) {
	te := newTestEnv(t)
	tests := []struct {
		name     string
		request  []byte
		response []byte
		code     int
	}{
		{
			name: "params",
			request: []byte(`{
								"address":"https://practicum.yandex.ru/"
							}`),
			response: []byte(`{
								"result":"http://localhost:8080/"
							}`),
			code: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(tt.request))
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			cfg := &config.Config{
				BaseURL: "http://localhost:8080",
			}

			h := te.httpHandler.PostBody(cfg)
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.code, rec.Code)
			}
		})
	}
}
