package handlers

import (
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"testing"
)

type testEnv struct {
	inMemoryRepo repository.URLStorage
	httpHandler  *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	te.inMemoryRepo = repository.NewInMemory()
	te.httpHandler = NewHTTPHandler(
		te.inMemoryRepo,
	)
	return te
}
