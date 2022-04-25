package repository

type URLStorage interface {
	Add(id string, url string)
	Get(id string) (string, error)
}
