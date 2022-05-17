package repository

import "sync"

type Memory struct {
	mx    sync.Mutex
	cache map[string]string
}

func NewInMemory() URLStorage {
	return &Memory{
		cache: make(map[string]string),
	}
}

func (s *Memory) Add(id string, url string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.cache[id] = url
	return nil
}

func (s *Memory) Get(id string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	URL := s.cache[id]
	return URL, nil
}
