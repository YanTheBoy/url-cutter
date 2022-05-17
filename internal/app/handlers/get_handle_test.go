package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	te := newTestEnv(t)
	type want struct {
		code int
		body *string
	}
	tests := []struct {
		name  string
		value string
		want  want
	}{
		{
			name:  "get without param",
			value: "",
			want:  want{code: 307},
		},
		{
			name:  "get with invalid param",
			value: "11111",
			want:  want{code: 307},
		},
		{
			name:  "get with params",
			value: "8406f18f-95c5-4953-94c8-275460c36f70",
			want:  want{code: 307},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.value)

			err := te.inMemoryRepo.Add("8406f18f-95c5-4953-94c8-275460c36f70",
				"https://www.uuidgenerator.net/version4")
			if err != nil {
				return
			}

			h := te.httpHandler.Get()
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
			}
		})
	}
}
