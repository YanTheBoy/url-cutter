package repository

import "fmt"

type Storage struct {
	urlStorage map[string]string
}

func CreateURLStorage() *Storage {
	return &Storage{
		urlStorage: make(map[string]string),
	}
}

func (d *Storage) Add(id string, url string) {
	d.urlStorage[id] = url
}

func (d *Storage) Get(id string) (string, error) {
	URL, found := d.urlStorage[id]
	if !found {
		return "", fmt.Errorf("cant find URL %s", id)
	}
	return URL, nil
}
