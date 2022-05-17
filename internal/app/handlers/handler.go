package handlers

import (
	"github.com/iliarkhpv/url-cutter/internal/repository"
	"net/url"
)

type HTTPHandler struct {
	urlStorage repository.URLStorage
}

func NewHTTPHandler(urlStorage repository.URLStorage) *HTTPHandler {
	return &HTTPHandler{
		urlStorage: urlStorage,
	}
}

func checkURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}
	return nil
}
