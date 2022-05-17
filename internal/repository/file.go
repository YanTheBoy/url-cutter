package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type InFile struct {
	mx      sync.Mutex
	cache   map[string]string
	encoder *json.Encoder
}

func NewInFile(filePath string) (URLStorage, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	cache := make(map[string]string)
	if fileInfo, _ := file.Stat(); fileInfo.Size() != 0 {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			err = json.Unmarshal(scanner.Bytes(), &cache)
			if err != nil {
				return nil, err
			}
		}
	}

	return &InFile{
		cache:   cache,
		encoder: json.NewEncoder(file),
	}, nil
}

func (s *InFile) Add(id string, url string) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.cache[id] = url

	data := make(map[string]string, 1)
	data[id] = url
	return s.encoder.Encode(&data)
}

func (s *InFile) Get(id string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	URL := s.cache[id]
	return URL, nil
}
