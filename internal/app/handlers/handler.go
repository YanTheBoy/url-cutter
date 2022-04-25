package handlers

import "github.com/iliarkhpv/url-cutter/internal/repository"

const host = "http://localhost:8080/"

type HTTPHandler struct {
	urlStorage repository.URLStorage
}

func NewHTTPHandler(urlStorage repository.URLStorage) *HTTPHandler {
	return &HTTPHandler{
		urlStorage: urlStorage,
	}
}
