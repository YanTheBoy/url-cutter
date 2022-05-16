package repository

type URLStorage interface {
	Add(id string, url string) error
	Get(id string) (string, error)
}
