package repository

import (
	"fmt"
	"sync"
)

type Storage struct {
	mx         sync.Mutex
	urlStorage map[string]string
}

type StorageCreator interface {
	CreateURLStorage()
}

func (d *Storage) CreateURLStorage() *Storage {
	return &Storage{
		urlStorage: make(map[string]string),
	}
}

func (d *Storage) Add(id string, url string) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.urlStorage[id] = url
}

func (d *Storage) Get(id string) (string, error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	URL, found := d.urlStorage[id]
	if !found {
		return "", fmt.Errorf("cant find URL %s", id)
	}
	return URL, nil
}
