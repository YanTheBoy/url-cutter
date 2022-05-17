package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostURL(t *testing.T) {
	te := newTestEnv(t)
	type want struct {
		code int
	}
	tests := []struct {
		name  string
		value string
		want  want
	}{
		{
			name:  "post with empty body is",
			value: "",
			want:  want{code: 400},
		},
		{
			name:  "post with param",
			value: "8406f18f-95c5-4953-94c8-275460c36f70",
			want:  want{code: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.value))
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := te.httpHandler.Post()
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
				require.NotNil(t, rec.Body.String())
			}
		})
	}
}
