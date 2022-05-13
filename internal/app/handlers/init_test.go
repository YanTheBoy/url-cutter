package handlers

import (
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"testing"
)

type testEnv struct {
	urlStorage  repository.URLStorage
	httpHandler *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	te.urlStorage = repository.CreateUrlStorage()
	te.httpHandler = NewHTTPHandler(
		te.urlStorage,
	)
	return te
}
