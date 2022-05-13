package repository

import (
	"fmt"
	"sync"
)

type Storage struct {
	mx         sync.Mutex
	urlStorage map[string]string
}

func CreateURLStorage() URLStorage {
	return &Storage{
		urlStorage: make(map[string]string),
	}
}

func (s *Storage) Add(id string, url string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.urlStorage[id] = url
}

func (s *Storage) Get(id string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	URL, found := s.urlStorage[id]
	if !found {
		return "", fmt.Errorf("cant find URL %s", id)
	}
	return URL, nil
}
