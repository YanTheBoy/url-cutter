package handlers

import "github.com/iliarkhpv/url-cutter/internal/repository"

const host = "http://localhost:8080/"

type HTTPHandler struct {
	urlRepository repository.URLStorage
}

func NewHTTPHandler(urlRepository repository.URLStorage) *HTTPHandler {
	return &HTTPHandler{
		urlStorage: urlStorage,
	}
}
